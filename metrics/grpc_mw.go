package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// ServerMetricsUnaryInterceptor
// is a gRPC server-side interceptor that provides Prometheus monitoring for Unary RPCs.
//
// For every RPC it exports the following metrics:
// - server_grpc_request_count{method, code}
// - server_grpc_response_time{method}
// namespace here is a prefix for metrics name
func ServerMetricsUnaryInterceptor(namespace string) func(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	var serverRequestCounter = MustRegisterCounterVec("server_grpc_request_count",
		namespace,
		"server_request_count", []string{"method", "code"})

	var serverResponseTime = MustRegisterHistogramVec("server_grpc_response_time",
		namespace,
		"server response time in seconds",
		TimeBucketsMedium, []string{"method"})

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		h, err := handler(ctx, req)
		tookTime := float64(time.Since(startTime)) / float64(time.Second)
		serverResponseTime.WithLabelValues(info.FullMethod).Observe(tookTime)
		if err != nil {
			serverRequestCounter.WithLabelValues(info.FullMethod, "500").Inc()
			return h, err
		}
		serverRequestCounter.WithLabelValues(info.FullMethod, "200").Inc()
		return h, err
	}
}
