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

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Debugf(format, args...)
		return
	}
	log.Debugf(format, args...)
}

func Debug(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Debug(v...)
		return
	}
	log.Debug(v...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Errorf(format, args...)
		return
	}
	log.Errorf(format, args...)
}

func Error(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Error(v...)
		return
	}
	log.Error(v...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Warnf(format, args...)
		return
	}
	log.Warnf(format, args...)
}

func Warn(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Warn(v...)
		return
	}
	log.Warn(v...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Fatalf(format, args...)
		return
	}
	log.Fatalf(format, args...)
}

func Fatal(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Fatal(v...)
		return
	}
	log.Fatal(v...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Panicf(format, args...)
		return
	}
	log.Panicf(format, args...)
}

func Panic(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Panic(v...)
		return
	}
	log.Panic(v...)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Tracef(format, args...)
		return
	}
	log.Tracef(format, args...)
}

func Trace(ctx context.Context, v ...interface{}) {
	log, ok := GetEntryFromCTX(ctx)
	if !ok {
		logrus.Trace(v...)
		return
	}
	log.Trace(v...)
}
