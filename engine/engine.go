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
	log "github.com/sirupsen/logrus"
)

type engineCfg struct {
	DebugPort       string        `json:"debug_port" envconfig:"DEBUG_PORT" default:":8084"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
}

type ServerOpt func(*Server)

//WithDisableBanner disable witcher Banner
func WithDisableBanner() ServerOpt {
	return func(s *Server) {
		s.disableBanner = true
	}
}

//WithShutdownTimeout set shutdown timeout
//
// timeout is duration for graceful shutdown
func WithShutdownTimeout(timeout time.Duration) ServerOpt {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

//WithDebugPort set debug port
// use value with `:` for example: :8084
func WithDebugPort(port string) ServerOpt {
	return func(server *Server) {
		server.PORT = port
	}
}

//Server is app engine with debug server
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
}

//Run start your server
func (s *Server) Run() error {
	log.Infof("Starting ...")

	//add debug server
	s.AddActor(func() error {
		err := s.DebugServer.RunDebug()
		if err.Error() != "http: Server closed" {
			return err
		}
		return nil
	}, func(err error) {
		if err != nil {
			log.Fatalf("Error Run Debug Server: %s", err)
		}
	})

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
	log.Infof("Start shutdown server")
	err := s.Shutdown()
	if err != nil {
		log.Errorf("Error Shutdown Server: %s", err)
		return err
	}
	log.Infof("Shutdown server gracefully")
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
			log.Errorf("Error runGroup Server: %s", err)
			s.setErrRunGroup(err)
		}
	}()

	s.AddChecker(NewDefaultChecker("run group", func() error {
		return s.getErrRunGroup()
	}))
}

//AddActor add actor control you background task
// you have execute function and done function(interrupt function)
// interrupt function handle the error
// execute function is called when server is ready
func (s *Server) AddActor(execute func() error, interrupt func(err error)) {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debugf("Add New Actor")
	s.runGroup.Add(execute, interrupt)
}

//AddCloser add closer object for shutdown your application gracefully
// you need to handle the context and call wg.Done() when done
//
func (s *Server) AddCloser(closer Closer) {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debugf("AddCloser %s", closer.GetName())
	s.closerGroup = append(s.closerGroup, closer)
}

func (s *Server) AddClosers(closers []Closer) {
	for _, closer := range closers {
		s.AddCloser(closer)
	}
}

//Shutdown server
//
// Timeout is duration for graceful shutdown
// If all goroutines are finished, server will be shutdown
// or timeout will be reached, server will be shutdown force
func (s *Server) Shutdown() error {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debugf("Shutdown with timeout - %s", s.shutdownTimeout.String())
	s.SetReady(false)
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.cancelFunc()

	err := s.ShutdownDebug(ctx)
	if err != nil {
		log.Errorf("Error Shutdown Debug Server: %s", err)
	}

	quit := make(chan struct{})
	go func() {
		err = s.closeClosers(ctx)
		if err != nil {
			quit <- struct{}{}
			log.Errorf("Error Close Closers: %s", err)
			return
		}
		quit <- struct{}{}
	}()

	for {
		select {
		case <-quit:
			return err
		case <-ctx.Done():
			log.Debugf("force shutdown")
			return errors.New("force quit")
		}
	}
}

func (s *Server) closeClosers(ctx context.Context) error {
	var wg sync.WaitGroup
	var errMsg []string
	for _, closer := range s.closerGroup {
		log.Debugf("Close closer: %s", closer.GetName())
		wg.Add(1)
		go func(c Closer) {
			err := c.Close(ctx, &wg)
			if err != nil {
				errMsg = append(errMsg, fmt.Sprintf("Error Close Closer %s: %s", c.GetName(), err))
				log.Errorf("Error Close Closer %s: %s", c.GetName(), err)
			}
		}(closer)
	}
	wg.Wait()
	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, "\n"))
	}
	log.Debug("Close all closers gracefully")
	return nil
}

//GetCTX return context of your engine
func (s *Server) GetCTX() context.Context {
	s.m.Lock()
	defer s.m.Unlock()
	return s.ctx
}

//NewServer create new server
//
//You can change config engine use ENV variables DEBUG_PORT and SHUTDOWN_TIMEOUT(for example 30s)
// or use With function
func NewServer(opt ...ServerOpt) *Server {
	var cfg engineCfg

	ctx, cancel := context.WithCancel(context.Background())

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Errorf("err read engine config: %s", err)
		panic(err)
	}

	debug := NewDebugServer(cfg.DebugPort)

	s := Server{
		DebugServer:     debug,
		shutdownTimeout: cfg.ShutdownTimeout,
		ctx:             ctx,
		cancelFunc:      cancel,
		runGroup:        run.Group{},
		closerGroup:     make([]Closer, 0, 10),
	}

	for _, opt := range opt {
		opt(&s)
	}
	return &s
}
