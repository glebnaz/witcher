package main

import (
	"os"

	"github.com/glebnaz/witcher/log"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	log.InitLog(os.Stdout, true)

	zlog.Debug().Msg("MSG")
}

//func main() {
//	log.InitLog(os.Stdout, true)
//
//	zlog.Debug().Msg("MSG")
//}

//      OUTPUT
// {"level":"info","t":"2023-11-28T21:32:59+01:00","msg":"Inited prod version logger"}
// {"level":"debug","t":"2023-11-28T21:32:59+01:00","msg":"MSG"}
