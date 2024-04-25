package metrics

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPMiddleware is a middleware for the standard net/http package.
//
// It collects metrics for HTTP requests.
// It returns a function that takes a handler and returns a handler.
// This is a pattern for net/http middleware.
//
// For every request, it exports the following metrics:
// - server_http_request_count{method,path,code}
// - server_http_response_time{method,path}
func HTTPMiddleware(namespace string) func(http.Handler) http.Handler {
	var serverRequestCounter = MustRegisterCounterVec("server_http_request_count",
		namespace,
		"server_request_count", []string{"method", "path", "code"})

	var serverResponseTime = MustRegisterHistogramVec("server_http_response_time",
		namespace,
		"server response time in seconds",
		TimeBucketsMedium, []string{"method", "path"})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			statusCode := fmt.Sprintf("%d", http.StatusOK) // replace with actual status code if available
			tookTime := float64(time.Since(startTime)) / float64(time.Second)
			serverResponseTime.WithLabelValues(r.Method, r.RequestURI).Observe(tookTime)
			serverRequestCounter.WithLabelValues(r.Method, r.RequestURI, statusCode).Inc()
		})
	}
}
