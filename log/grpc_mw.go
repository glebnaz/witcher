package log

import (
	"context"

	"github.com/glebnaz/witcher/trace"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServerLoggerUnaryInterceptor() func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, reqID := trace.MustGetSimpleReqIDFromContext(ctx)
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
