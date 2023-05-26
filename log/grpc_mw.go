package log

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	unknown = "unknown"
)

func ServerLoggerUnaryInterceptor() func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqID := unknown
		fields := logrus.Fields{
			"method": info.FullMethod,
			"req_id": reqID,
		}
		log := logrus.WithFields(fields)
		ctx = AddEntryToCTX(ctx, log)
		ctx = AddFieldsToCTX(ctx, fields)
		return handler(ctx, req)
	}
}
