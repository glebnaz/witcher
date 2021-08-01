package logger

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	loggerKey = "logger"
	unknown   = "unknown"
)

func GenerateLoggerHTTP(c echo.Context) *logrus.Entry {
	l := logrus.WithFields(map[string]interface{}{
		"PATH":   c.Path(),
		"METHOD": c.Request().Method,
		"IP":     c.RealIP(),
		"UA":     c.Request().UserAgent(),
	})
	return l
}

//HTTPFromContext return http logger from context
//logger is logrus *Entry
func HTTPFromContext(ctx context.Context) *logrus.Entry {
	l, ok := ctx.Value(loggerKey).(*logrus.Entry)
	if !ok {
		return logrus.WithFields(map[string]interface{}{
			"PATH":   unknown,
			"METHOD": unknown,
			"IP":     unknown,
			"UA":     unknown,
		})
	}
	return l
}
