package grpc

import (
	"context"
	"github.com/glebnaz/go-platform/logger"

	"google.golang.org/grpc"
)

//NewServerUnaryLoggerInterceptor return logger Unary Interceptor
func NewServerUnaryLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		l := logger.GenerateLoggerUnaryGRPC(ctx, info)
		l.Debug("New request")
		ctx = logger.ToContext(ctx, l)
		return handler(ctx, req)
	}
}
