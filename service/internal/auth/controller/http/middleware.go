package http

import (
	"net/http"
	"service/internal/auth/metrics"
	"time"
)

func (eh *EndpointHandler) metricsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		path := r.URL.Path

		statusString := rw.GetStatusString()

		metrics.HttpResponseTime.WithLabelValues(path, statusString, r.Method).Observe(time.Since(start).Seconds())
		metrics.HttpRequestsTotalCollector.WithLabelValues(path, statusString, r.Method).Inc()
	})
}
