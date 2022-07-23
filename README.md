[![Sourcegraph](https://sourcegraph.com/github.com/glebnaz/witcher/-/badge.svg?style=flat-square)]

# WITCHER

Golang web-app engine. Completely ready to work in a cluster!



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
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⢿⣿⣿⡿⠿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀


## Quckstart

```

```go
package main

import (
	"context"
	"sync"
	"time"

	"github.com/glebnaz/witcher/engine"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})

	s := engine.NewServer()

	checker := engine.NewDefaultChecker("checker", func() error {
		log.Infof("Checker is running")
		return nil
	})

	closer := engine.NewDefaultCloser("closer", func(ctx context.Context, wg *sync.WaitGroup) error {
		log.Infof("Closer is running")
		time.Sleep(time.Second * 10)
		log.Infof("Closer is running after 10 seconds")
		wg.Done()
		return nil
	})

	s.AddChecker(checker)
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

	if err := s.Run(); err != nil {
		log.Errorf("Error Run Server: %s", err)
	}
}
```

## Debug Server

### Handlers
In your Debug server you have tree handlers: 

/live - live probe 
/ready - ready probe
/metrics - metrics prometheus type

By default, the server is listening on port 8084. You can change it by setting the environment variable `PORT`.
Or you can set port by WithDebugPort("your_port)

### Check function
You can add checkers and closers by AddChecker and AddCloser. And add actors by AddActor.
Checkers use for live probe, Actor is for goroutine control, Closer is for graceful shutdown.

## Metrics
Use metrics prometheus type. Use metrics pkg to register metrics.


