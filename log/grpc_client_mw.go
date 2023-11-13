package log

import (
	"context"

	"google.golang.org/grpc"
)

// ClientLoggerUnaryInterceptor returns a new unary client interceptor for
// log from what to
func ClientLoggerUnaryInterceptor(from, to string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		/////////////////////////////////////////////////

		Debugf(ctx, "Call %s from %s to %s", method, from, to)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
