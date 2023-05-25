package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

func Infof(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Infof(format, args...)
		return
	}
	log.Infof(format, args...)
}

func Info(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Info(v...)
		return
	}
	log.Info(v...)
}
