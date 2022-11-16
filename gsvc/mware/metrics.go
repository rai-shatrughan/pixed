package mware

import (
	"net/http"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

func InitMetricMiddleware() middleware.Middleware {
	return middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{
			Prefix: serviceName,
		}),
	})
}

func PrometheusMiddleware(mdlw middleware.Middleware, next http.Handler) http.Handler {
	return std.Handler("", mdlw, next)
}
