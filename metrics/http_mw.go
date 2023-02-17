package metrics

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

// EchoMiddleware is a middleware for echo framework
// It collects metrics for http requests
// It returns a function that takes a handler and returns a handler/
// This is a pattern for echo middleware
// https://echo.labstack.com/guide/middleware
// https://echo.labstack.com/cookbook/middleware
// https://echo.labstack.com/cookbook/middleware#custom-middleware
//
//
//// For every RPC it exports the following metrics:
//// - server_http_request_count{method, code}
//// - server_http_response_time{method}
func EchoMiddleware(namespace string) echo.MiddlewareFunc {
	var serverRequestCounter = MustRegisterCounterVec("server_http_request_count",
		namespace,
		"server_request_count", []string{"method", "code"})

	var serverResponseTime = MustRegisterHistogramVec("server_http_response_time",
		namespace,
		"server response time in seconds",
		TimeBucketsMedium, []string{"method"})
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now()
			err := next(c)
			statusCode := fmt.Sprintf("%d", c.Response().Status)
			tookTime := float64(time.Since(startTime)) / float64(time.Second)
			serverResponseTime.WithLabelValues(c.Request().Method).Observe(tookTime)
			serverRequestCounter.WithLabelValues(c.Request().Method, statusCode).Inc()
			return err
		}
	}
}
