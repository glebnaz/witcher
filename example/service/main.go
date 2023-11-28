package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/glebnaz/witcher/db/mongo"
	"github.com/glebnaz/witcher/engine"
	"github.com/glebnaz/witcher/grpc"
	wlog "github.com/glebnaz/witcher/log"
	"github.com/rs/zerolog/log"
)

func main() {
	wlog.InitLog(os.Stdout, false)
	serverGRPC := grpc.NewDefaultServer("witcher_example")
	s := engine.NewServer(engine.WithGRPCServer(serverGRPC, ":8082", time.Second*5))

	checker1 := engine.NewDefaultChecker("checker false", func() error {
		log.Info().Msgf("Checker1 is running")
		return nil
	})

	checker2 := engine.NewDefaultChecker("checker true", func() error {
		log.Info().Msgf("Checker2 is running")
		return nil
	})

	closer := engine.NewDefaultCloser("closer", func(ctx context.Context, wg *sync.WaitGroup) error {
		log.Info().Msgf("Closer is running")
		time.Sleep(time.Second * 10)
		log.Info().Msgf("Closer is running after 10 seconds")
		wg.Done()
		return nil
	})

	s.AddCheckers([]engine.Checker{checker1, checker2})
	s.AddCloser(closer)
	s.AddActor(func() error {
		i := 0
		for {
			i++
			log.Debug().Msgf("Actor is running %d", i)
			time.Sleep(time.Second * 1)
		}
	},
		func(error) {
			log.Error().Msg("Actor failed")
		})

	m, err := mongo.NewMongo(s.GetCTX(), "mongodb://localhost:27017")
	if err != nil {
		log.Error().Msgf("Error connect to mongo: %s", err)
	}

	s.AddCloser(m.Closer())
	s.AddChecker(m.HealthChecker(time.Second * 5))

	if err := s.Run(); err != nil {
		log.Error().Msgf("Error Run Server: %s", err)
	}
}
