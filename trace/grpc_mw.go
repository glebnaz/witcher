package trace

import (
	"context"

	"google.golang.org/grpc"
)

// ServerSimpleRequestIDUnaryInterceptor returns a new unary server interceptor for
// simple request ID
// if reqID is not found in the context, generates a new one and adds it to the context
func ServerSimpleRequestIDUnaryInterceptor() func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		/////////////////////////////////////////////////
		reqID := GetSimpleReqIDFromContext(ctx)
		if reqID == "" {
			reqID = GenerateSimpleReqID()
			ctx = AddSimpleReqIDToContext(ctx, reqID)
		}
		return handler(ctx, req)
	}
}
