package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var metricsCollector = make(map[string]any)

// ClientMetricsUnaryInterceptor returns a new unary client interceptor for
// metrics from what to
// namespace here is a prefix for metrics name
//
// For every RPC it exports the following metrics:
// - client_grpc_request_count{method, code, from, to}
// - client_grpc_response_time{method, from, to}
// namespace here is a prefix for metrics name
func ClientMetricsUnaryInterceptor(from, to string) grpc.UnaryClientInterceptor {
	const nameCounter = "client_grpc_request_count"
	const nameHistogram = "client_grpc_response_time"

	keyCounter := nameCounter + from
	keyHistogram := nameHistogram + from

	var clientRequestCounter *prometheus.CounterVec
	var clientResponseTime *prometheus.HistogramVec

	if _, ok := metricsCollector[keyCounter]; !ok {
		clientRequestCounter = MustRegisterCounterVec("client_grpc_request_count",
			from,
			"client_request_count", []string{"method", "code", "from", "to"})

		metricsCollector[keyCounter] = clientRequestCounter
	} else {
		clientRequestCounter = metricsCollector[keyCounter].(*prometheus.CounterVec)
	}

	if _, ok := metricsCollector[keyHistogram]; !ok {
		clientResponseTime = MustRegisterHistogramVec("client_grpc_response_time",
			from,
			"client response time in seconds",
			TimeBucketsMedium, []string{"method", "from", "to"})
		metricsCollector[keyHistogram] = clientResponseTime
	} else {
		clientResponseTime = metricsCollector[keyHistogram].(*prometheus.HistogramVec)
	}

	return func(ctx context.Context,
		method string,
		req,
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {

		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		tookTime := float64(time.Since(startTime)) / float64(time.Second)
		hStatus, ok := status.FromError(err)
		statusString := unknownCode
		if ok {
			statusString = hStatus.Code().String()
		}
		clientResponseTime.WithLabelValues(method, from, to).Observe(tookTime)
		clientRequestCounter.WithLabelValues(method, statusString, from, to).Inc()

		return err
	}
}
