package log

import (
	"io"
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

// InitLog initializes logrus logger
// It uses nested-logrus-formatter
// It sets log level from LOGLVL env variable
// data format: 2006|01|02 15:04:05.000
func InitLog(out io.Writer) {
	log.SetFormatter(&formatter.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006|01|02 15:04:05.000",
	})
	log.SetOutput(out)
	level, err := log.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil {
		level = log.DebugLevel
	}
	log.SetLevel(level)
}
