package metrics

import (
	"context"
	"google.golang.org/grpc"
)

// ServerMetricsUnaryInterceptor
//is a gRPC server-side interceptor that provides Prometheus monitoring for Unary RPCs.
//
// For every RPC it exports the following metrics:
// todo add metrics list
//namespace here is a prefix for metrics name
func ServerMetricsUnaryInterceptor(namespace string) func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	var serverRequestCounter = MustRegisterCounterVec("server_grpc_request_count",
		namespace,
		"server_request_count", []string{"method", "code"})

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		h, err := handler(ctx, req)
		if err != nil {
			serverRequestCounter.WithLabelValues(info.FullMethod, "500").Inc()
			return h, err
		}
		serverRequestCounter.WithLabelValues(info.FullMethod, "200").Inc()
		return h, err
	}
}
