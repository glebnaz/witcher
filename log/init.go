package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	TimestampFieldNameDefault = "t"
	MessageFieldNameDefault   = "msg"
)

// InitLog initializes zerolog for this service
// you need provide stdout for your log
// this function will init default logger with default formatter
// and default hook
func InitLog(out io.Writer, isProd bool) {
	if isProd {
		SetDefaultProductionLogger(out)
		log.Info().Msg("Inited prod version logger")
		return
	}
	SetDefaultDevLogger(out)
	log.Info().Msg("Inited dev version logger")
	return
}

func SetDefaultProductionLogger(w io.Writer) {
	setDefaultLogger(w)
}

func SetDefaultDevLogger(w io.Writer) {
	output := zerolog.ConsoleWriter{
		Out:        w,
		TimeFormat: "2006|01|02 15:04:05.000",
		NoColor:    false,
	}

	setDefaultLogger(output)
}

func setDefaultLogger(out io.Writer) {
	zerolog.TimestampFieldName = TimestampFieldNameDefault
	zerolog.MessageFieldName = MessageFieldNameDefault

	logger := zerolog.New(out).With().Timestamp().Logger()

	log.Logger = logger

	lvl, err := zerolog.ParseLevel(os.Getenv("LOGLVL"))
	if err != nil || lvl == zerolog.NoLevel {
		lvl = zerolog.DebugLevel
	}

	log.Logger = log.Level(lvl)

	//hooks
	for i := range defaultHooks {
		log.Logger = log.Hook(defaultHooks[i])
	}
}
