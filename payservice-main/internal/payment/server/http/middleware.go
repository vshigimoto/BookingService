package http

import (
	"fmt"
	"net/http"
	"payservice/internal/payment/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *EndpointHandler) TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")

		user, err := h.paymentService.VerifyToken(accessToken)
		fmt.Println("User Id is", user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "Auth token is invalid")
			return
		} else {
			ctx.Set("user", user)
		}

	}
}

func (h *EndpointHandler) MetricsHandler() gin.HandlerFunc {
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
