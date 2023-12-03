package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payservice/internal/auth/metrics"
	"time"
)

func (eh *EndpointHandler) MetricsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		rw := &responseWriter{c.Writer, http.StatusOK}

		c.Next()

		path := c.Request.URL.Path
		statusString := rw.GetStatusString()

		metrics.HttpResponseTime.WithLabelValues(path, statusString, c.Request.Method).Observe(time.Since(start).Seconds())
		metrics.HttpRequestsTotalCollector.WithLabelValues(path, statusString, c.Request.Method).Inc()
	}
}
