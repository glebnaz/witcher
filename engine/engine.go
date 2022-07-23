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

//todo add http and grpc server

//Server is app engine with debug server
type Server struct {
	*DebugServer

	ctx             context.Context
	cancelFunc      context.CancelFunc
	closerGroup     []Closer
	shutdownTimeout time.Duration
	m               sync.Mutex

	runGroup      run.Group
	errorRunGroup error
}

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
			log.Errorf("Error Run Debug Server: %s", err)
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

func (s *Server) introduce() {
	log.Infof("Server Is Ready")
	log.Infof(introduce, s.PORT)
	log.Infof(logo)
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

func (s *Server) AddActor(execute func() error, interrupt func(err error)) {
	s.m.Lock()
	defer s.m.Unlock()
	log.Debugf("Add New Actor")
	s.runGroup.Add(execute, interrupt)
}

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

func (s *Server) GetCTX() context.Context {
	s.m.Lock()
	defer s.m.Unlock()
	return s.ctx
}

func NewServer() *Server {
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
	return &s
}
