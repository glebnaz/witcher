package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Log struct {
	logrus.Logger
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log, ok := MustGetEntryFromCTX(ctx)
	if !ok {
		logrus.Infof(format, args...)
		return
	}
	log.Infof(format, args...)
}
