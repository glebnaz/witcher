package grpc

import (
	"github.com/glebnaz/witcher/metrics"
	"github.com/glebnaz/witcher/trace"
	"google.golang.org/grpc"
)

var defaultChainUnaryServerInterceptor = []grpc.UnaryServerInterceptor{
	trace.ServerSimpleRequestIDUnaryInterceptor(),
}

func NewDefaultServer(metricNamespace string, opt ...grpc.ServerOption) *grpc.Server {
	metricInterceptor := metrics.ServerMetricsUnaryInterceptor(metricNamespace)

	interceptors := []grpc.UnaryServerInterceptor{
		metricInterceptor,
	}

	interceptors = append(interceptors, defaultChainUnaryServerInterceptor...)

	newOPT := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptors...),
	}
	newOPT = append(newOPT, opt...)

	return grpc.NewServer(
		newOPT...,
	)
}
