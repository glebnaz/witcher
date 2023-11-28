package main

import (
	"os"

	"github.com/glebnaz/witcher/log"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	log.InitLog(os.Stdout, false)

	zlog.Debug().Msg("sds")
}
