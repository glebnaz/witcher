package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type engineCfg struct {
	DebugPort       string        `json:"debug_port" envconfig:"DEBUG_PORT" default:":8084"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

type ServerOpt func(*Server)

// WithDisableBanner disable witcher Banner
func WithDisableBanner() ServerOpt {
	return func(s *Server) {
		s.disableBanner = true
	}
}

// WithShutdownTimeout set shutdown timeout
//
//	is duration for graceful shutdown
func WithShutdownTimeout(timeout time.Duration) ServerOpt {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// WithDebugPort set debug port
// use value with `:` for example: :8084
func WithDebugPort(port string) ServerOpt {
	return func(server *Server) {
		server.PORT = port
	}
}

func WithGRPCServer(grpcServer *grpc.Server, port string, shutdownTimeout time.Duration) ServerOpt {
	return func(s *Server) {
		runner := newGRPCServerRunner(grpcServer, port, shutdownTimeout)
		s.grpcServerRunner = runner
	}
}

// Server is app engine with debug server
//
// Use AddCloser to add closer objects for shutdown your application gracefully(for more information see AddCloser)
//
// Use AddActor to add actor control you background task(for more information see AddActor)
//
// Use AddChecker to add checker for live probe(for more information see AddChecker)
//
// You can change config engine use ENV variables DEBUG_PORT and SHUTDOWN_TIMEOUT(for example 30s)
type Server struct {
	*DebugServer

	ctx             context.Context
	cancelFunc      context.CancelFunc
	closerGroup     []Closer
	shutdownTimeout time.Duration
	m               sync.Mutex

	runGroup      run.Group
	errorRunGroup error

	//options
	disableBanner bool

	grpcServerRunner *grpcServerRunner
}

// Run start your server
func (s *Server) Run() error {
	log.Info().Msgf("Starting ...")

	//add debug server
	s.AddActor(func() error {
		err := s.DebugServer.RunDebug()
		if err.Error() != "http: Server closed" {
			return err
		}
		return nil
	}, func(err error) {
		if err != nil {
			log.Fatal().Msgf("Error Run Debug Server: %s", err)
		}
	})

	if s.grpcServerRunner != nil {
		s.AddActor(s.grpcServerRunner.actor())
	}

	s.runRunGroup()

	//all ready
	s.SetReady(true)
	s.introduce()

	//---- shutdown ----

	//wait signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//start shutdown
	log.Info().Msgf("Start shutdown server")
	err := s.Shutdown()
	if err != nil {
		log.Error().Msgf("Error Shutdown Server: %s", err)
		return err
	}
	log.Info().Msgf("Shutdown server gracefully")
	return nil
}

func (s *Server) getErrRunGroup() error {
	s.m.Lock()
	defer s.m.Unlock()
	return s.errorRunGroup
}

func (s *Server) setErrRunGroup(err error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.errorRunGroup = err
}

func (s *Server) runRunGroup() {
	go func() {
		err := s.runGroup.Run()
		if err != nil {
			log.Error().Msgf("Error runGroup Server: %s", err)
			s.setErrRunGroup(err)
		}
	}()

	s.AddChecker(NewDefaultChecker("run group", func(_ context.Context) error {
		return s.getErrRunGroup()
	}))
}

// AddActor add actor control you background task
//
// you have executed function and done function(interrupt function)
// interrupt function handle the error
// execute function is called when server is ready
func (s *Server) AddActor(execute func() error, interrupt func(err error)) {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debug().Msgf("Add New Actor")
	s.runGroup.Add(execute, interrupt)
}

// AddCloser add closer object for shutdown your application gracefully
// you need to handle the context and call wg.Done() when done
func (s *Server) AddCloser(closer Closer) {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debug().Msgf("AddCloser %s", closer.GetName())
	s.closerGroup = append(s.closerGroup, closer)
}

func (s *Server) AddClosers(closers []Closer) {
	for _, closer := range closers {
		s.AddCloser(closer)
	}
}

// Shutdown server
//
// Timeout is duration for graceful shutdown
// If all goroutines are finished, server will be shutdown
// or timeout will be reached, server will be shutdown force
func (s *Server) Shutdown() error {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debug().Msgf("Shutdown with timeout - %s", s.shutdownTimeout.String())
	s.SetReady(false)
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.cancelFunc()

	err := s.ShutdownDebug(ctx)
	if err != nil {
		log.Error().Msgf("Error Shutdown Debug Server: %s", err)
	}

	quit := make(chan struct{})
	go func() {
		err = s.closeClosers(ctx)
		if err != nil {
			quit <- struct{}{}
			log.Error().Msgf("Error Close Closers: %s", err)
			return
		}
		quit <- struct{}{}
	}()

	for {
		select {
		case <-quit:
			return err
		case <-ctx.Done():
			log.Debug().Msgf("force shutdown")
			return errors.New("force quit")
		}
	}
}

type ErrorCloser struct {
	data map[string]error
}

func (e *ErrorCloser) Error() string {
	var err []string
	for k, v := range e.data {
		err = append(err, fmt.Sprintf("Error Close %s: %s", k, v))
	}
	return strings.Join(err, "\n")
}

func (s *Server) closeClosers(ctx context.Context) error {
	var wg sync.WaitGroup
	errCloser := ErrorCloser{data: make(map[string]error)}

	for _, closer := range s.closerGroup {
		log.Debug().Msgf("Close closer: %s", closer.GetName())
		wg.Add(1)
		go func(c Closer, wg *sync.WaitGroup) {
			defer wg.Done()
			err := c.Close(ctx)
			if err != nil {
				errCloser.data[c.GetName()] = err
				log.Error().Msgf("Error Close Closer %s: %s", c.GetName(), err)
			}
		}(closer, &wg)
	}
	wg.Wait()
	if len(errCloser.data) > 0 {
		return &errCloser
	}
	log.Debug().Msg("Close all closers gracefully")
	return nil
}

// GetCTX return context of your engine
func (s *Server) GetCTX() context.Context {
	s.m.Lock()
	defer s.m.Unlock()
	return s.ctx
}

// NewServer create new server
//
// You can change config engine use ENV variables DEBUG_PORT and SHUTDOWN_TIMEOUT(for example 30s)
// or use With function
func NewServer(opt ...ServerOpt) *Server {
	var cfg engineCfg

	ctx, cancel := context.WithCancel(context.Background())

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Error().Msgf("err read engine config: %s", err)
		panic(err)
	}

	debug := NewDebugServer(cfg.DebugPort)

	s := Server{
		DebugServer:      debug,
		shutdownTimeout:  cfg.ShutdownTimeout,
		ctx:              ctx,
		cancelFunc:       cancel,
		runGroup:         run.Group{},
		closerGroup:      make([]Closer, 0, 10),
		grpcServerRunner: nil,
	}

	for _, opt := range opt {
		opt(&s)
	}
	return &s
}
