package http

import (
	"github.com/glebnaz/go-platform/logger"
	"github.com/glebnaz/go-platform/metrics"
	"github.com/labstack/echo"
)

var httpMetric = metrics.MustRegisterCounterVec("http_method", "default",
	"default metric for http, count http request", []string{"success", "method", "path"})

func Logger(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GenerateLoggerHTTP(c)
		log.Debug("new request")
		err := handlerFunc(c)
		if err != nil {
			httpMetric.WithLabelValues("0", c.Request().Method, c.Path()).Inc()
			log.Error("request with error")
		} else {
			httpMetric.WithLabelValues("1", c.Request().Method, c.Path()).Inc()
			log.Debug("request result success")
		}
		return err
	}
}
