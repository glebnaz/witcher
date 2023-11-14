package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

const (
	unknownCode = "unknown"
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

	const nameCounter = "server_grpc_request_count"
	const nameHistogram = "server_grpc_response_time"

	keyCounter := nameCounter + namespace
	keyHistogram := nameHistogram + namespace

	var serverRequestCounter *prometheus.CounterVec
	var serverResponseTime *prometheus.HistogramVec

	if _, ok := metricsCollector[keyCounter]; !ok {
		serverRequestCounter = MustRegisterCounterVec(nameCounter,
			namespace,
			"server_request_count", []string{"method", "code"})
		metricsCollector[keyCounter] = serverRequestCounter
	} else {
		serverRequestCounter = metricsCollector[keyCounter].(*prometheus.CounterVec)
	}

	if _, ok := metricsCollector[keyHistogram]; !ok {
		serverResponseTime = MustRegisterHistogramVec(nameHistogram,
			namespace,
			"server response time in seconds",
			TimeBucketsMedium, []string{"method"})
		metricsCollector[keyHistogram] = serverResponseTime
	} else {
		serverResponseTime = metricsCollector[keyHistogram].(*prometheus.HistogramVec)
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		h, err := handler(ctx, req)
		tookTime := float64(time.Since(startTime)) / float64(time.Second)
		hStatus, ok := status.FromError(err)
		statusString := unknownCode
		if ok {
			statusString = hStatus.Code().String()
		}
		serverResponseTime.WithLabelValues(info.FullMethod).Observe(tookTime)
		serverRequestCounter.WithLabelValues(info.FullMethod, statusString).Inc()
		return h, err
	}
}
