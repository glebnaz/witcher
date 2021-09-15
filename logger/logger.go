package logger

import (
	"context"

	"github.com/glebnaz/go-platform/trace"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	loggerKey = "logger"
	unknown   = "unknown"
)

//GenerateLoggerHTTP generate logger from echo.Context
func GenerateLoggerHTTP(c echo.Context) *logrus.Entry {
	reqID := trace.GetReqIDFromCTX(c.Request().Context())
	l := logrus.WithFields(map[string]interface{}{
		"PATH":   c.Path(),
		"METHOD": c.Request().Method,
		"IP":     c.RealIP(),
		"UA":     c.Request().UserAgent(),
		"REQ-ID": reqID,
	})
	return l
}

//FromContext return logger from context
//logger is logrus *Entry
func FromContext(ctx context.Context) *logrus.Entry {
	l, ok := ctx.Value(loggerKey).(*logrus.Entry)
	if !ok {
		reqID := trace.GetReqIDFromCTX(ctx)
		return logrus.WithFields(map[string]interface{}{
			"PATH":   unknown,
			"METHOD": unknown,
			"IP":     unknown,
			"UA":     unknown,
			"REQ-ID": reqID,
		})
	}
	return l
}

//ToContext add logger to ctx
//return logger
func ToContext(ctx context.Context, l *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

//GenerateLoggerUnaryGRPC generate logger from grpc Unary
func GenerateLoggerUnaryGRPC(ctx context.Context, info *grpc.UnaryServerInfo) *logrus.Entry {
	reqID := trace.GetReqIDFromCTX(ctx)
	l := logrus.WithFields(map[string]interface{}{
		"PATH":   info.FullMethod,
		"METHOD": info.FullMethod,
		"IP":     unknown,
		"UA":     unknown,
		"REQ-ID": reqID,
	})
	return l
}
