# WITCHER
![Sourcegraph](https://sourcegraph.com/github.com/glebnaz/witcher/-/badge.svg?style=flat-square)
![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)
[![Release](https://img.shields.io/github/release/glebnaz/witcher.svg?style=flat-square)](https://github.com/glebnaz/witcher/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/glebnaz/witcher)](https://goreportcard.com/report/github.com/glebnaz/witcher)


⠀⠀⠀⠀⠀⠀⢠⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣳⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣿⣿⣶⣄⣀⠀⠀⠀⠀⠀⠀⠈⢳⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⡞⠁⠀⠀⠀⠀⠀⠀⣀⣠⣶⣿⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠘⠷⣽⣻⢿⣿⣶⣤⣄⡀⠀⠀⠀⠹⣿⣷⣄⢶⣄⠹⣿⣿⣿⣿⣿⣿⣿⣿⠏⣠⡶⣠⣾⣿⠏⠀⠀⠀⢀⣠⣤⣶⣿⡿⣟⣯⡾⠃⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢱⠈⠛⢷⣮⣟⠿⣿⣿⣿⣶⣦⣄⡘⢿⣿⡎⣿⣷⡹⣿⣿⣿⣿⣿⣿⢏⣾⡿⢱⣿⡿⢃⣠⣴⣶⣿⣿⣿⠿⣻⣽⡾⠛⠁⡎⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢣⢳⣄⡉⠻⢿⣶⣝⡻⢿⣿⢟⣼⣷⣿⣿⣿⣿⣷⣽⣿⡿⢿⣿⣯⣾⣿⣿⣿⣿⣿⣧⡻⣿⡿⢟⣫⣶⡿⠟⢉⣠⡞⡼⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⢧⡹⣿⣦⣀⠈⠛⢿⡷⣠⣛⠿⣿⣿⣏⣍⢿⣿⣿⡇⣿⣿⢸⣿⣾⡿⣩⣹⣿⣿⠿⣛⣄⢾⡿⠛⠁⣀⣴⣿⢏⡼⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠳⡜⢿⣿⣷⡄⠀⢘⠻⢿⣿⣮⣝⡺⢿⣷⡽⣿⣿⢻⡟⣿⣿⢯⣾⡿⢗⣫⣵⣿⡿⠟⡃⠀⢠⣾⣿⡿⢣⠞⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠈⠉⠛⠻⠿⢿⣆⠙⢦⣙⠿⢀⣰⣿⣷⡀⠈⣝⡻⢿⣷⣯⣓⢹⣿⢸⡇⣿⡏⣚⣽⣾⡿⢟⣫⠁⢀⣾⣿⣆⡀⠿⣋⡴⠋⣰⡿⠿⠟⠛⠉⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣄⡺⢿⣿⣿⣟⣷⣄⠘⠛⠛⠈⠛⠻⢼⣿⡟⢳⣿⡧⠿⠛⠁⠛⠛⠃⣀⣾⣹⣿⣿⡿⢗⣠⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣾⣿⣿⣿⣿⣷⣬⣛⠿⣮⣝⣻⣿⣓⡍⣷⢸⣿⣿⡇⢸⣿⣿⡇⣶⢩⣚⣿⣟⣫⣵⠿⣛⣥⣶⣿⣿⣿⣿⣷⣶⣤⣀⡀⠀⠀⠀⠀⠀⠀
⠀⢀⣀⣤⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣝⠻⢿⣿⣿⡿⣼⣿⣿⣿⣿⣿⣿⣧⢿⣿⣿⡿⠟⣫⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣤⣀⡀⠀
⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠉⢻⣝⡻⢿⣿⣿⣿⣿⡿⢟⣫⡟⠉⠀⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣾⣿⣿⣿⣿⡶⠄⠸⣿⣿⣄⡀⠀⠀⢀⣠⣿⣿⡇⠠⢶⣿⣿⣿⣿⣷⣦⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⣿⣿⣿⣿⡿⠟⠉⠀⠀⠀⠛⠛⠛⠛⢛⡛⠛⠛⠛⠛⠀⠀⠀⠉⠻⢿⣿⣿⣿⣿⣿⣿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⢸⣧⢠⡀⠀⠀⢀⡄⣼⡇⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣀⡀⠀⢸⠃⠸⠁⠏⠹⠈⠇⠘⡇⠀⢀⣀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣀⣴⣿⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣦⣀⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣾⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⠏⠀⠀⠀⣿⡄⢀⠀⠀⠀⠀⡀⢠⣿⠀⠀⠀⠹⣿⣿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣷⡄⠀⠀⠀⠀
⠀⠀⠀⣰⣿⠟⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⠃⠀⠀⠀⠀⢹⠇⢸⣷⣄⣠⣾⡇⠸⡟⠀⠀⠀⠀⠘⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠻⣿⣆⠀⠀⠀
⠀⠀⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⡿⠁⠀⠀⠀⠀⠀⢸⢸⢸⣿⣿⣿⣿⡇⡇⡇⠀⠀⠀⠀⠀⠈⢿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠳⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⡟⠀⠀⠀⠀⠀⠀⠀⢸⡜⠰⣿⣿⣿⣿⠆⢣⡇⠀⠀⠀⠀⠀⠀⠀⢻⣿⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠈⠁⠈⠁⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⡄⠀⠀⠀⠀⢠⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢃⡷⢠⣤⣤⡄⢾⡸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢨⣥⣬⣉⣉⣥⣬⡅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⢿⣿⣿⡿⠿⠃⠀⠀⠀⠀⠀⠀⠀⠀

**Golang web-app engine. Completely ready to work in a cluster!**

**Scale your web app effortlessly with our Golang engine built for clusters**

# Documentation

## Quckstart

```go
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

```

## Debug Server

### Handlers
Your Debug server has three handlers:

* /live - live probe

* /read - readiness probe

* /startup - startup 

* /metrics - Prometheus metrics

#### PPROF
You can use pprof to see the CPU and memory usage of your application on the Debug server. Go to "/debug/pprof/" to see the start page of pprof.

By default, the server is listening on port 8084. You can change the port by setting the environment variable `PORT` or by using `WithDebugPort("your_port")`.

### Check Function
You can add checkers, closers, and actors to your server by using the `AddChecker`, `AddCloser`, and `AddActor` functions, respectively.

Checkers are used for live probes, actors are used for goroutine control, and closers are used for graceful shutdown.

## Metrics
You can use the Prometheus metrics type with the `metrics` package to register metrics.

## More Examples
You can find more examples in the `example` folder.

