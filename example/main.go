package main

import (
	"context"
	"sync"
	"time"

	"github.com/glebnaz/witcher/db/mongo"
	"github.com/glebnaz/witcher/engine"
	"github.com/glebnaz/witcher/metrics"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	serverGRPC := grpc.NewServer(grpc.UnaryInterceptor(metrics.ServerMetricsUnaryInterceptor("test")))
	s := engine.NewServer(engine.WithGRPCServer(serverGRPC, ":8082", time.Second*5))

	checker1 := engine.NewDefaultChecker("checker false", func() error {
		log.Infof("Checker1 is running")
		return nil
	})

	checker2 := engine.NewDefaultChecker("checker true", func() error {
		log.Infof("Checker2 is running")
		return nil
	})

	closer := engine.NewDefaultCloser("closer", func(ctx context.Context, wg *sync.WaitGroup) error {
		log.Infof("Closer is running")
		time.Sleep(time.Second * 10)
		log.Infof("Closer is running after 10 seconds")
		wg.Done()
		return nil
	})

	s.AddCheckers([]engine.Checker{checker1, checker2})
	s.AddCloser(closer)
	s.AddActor(func() error {
		i := 0
		for {
			i++
			log.Debugf("Actor is running %d", i)
			time.Sleep(time.Second * 1)
		}
	},
		func(error) {
			log.Errorf("Actor failed")
		})

	m, err := mongo.NewMongo(s.GetCTX(), "mongodb://localhost:27017")
	if err != nil {
		log.Errorf("Error connect to mongo: %s", err)
	}

	s.AddCloser(m.Closer())
	s.AddChecker(m.HealthChecker(time.Second * 5))

	if err := s.Run(); err != nil {
		log.Errorf("Error Run Server: %s", err)
	}
}
