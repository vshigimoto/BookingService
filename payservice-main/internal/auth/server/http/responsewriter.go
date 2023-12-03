package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type responseWriter struct {
	gin.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) GetStatusString() string {
	status := rw.status
	if rw.status == 0 {
		status = http.StatusOK
	}

	return strconv.Itoa(status)
}
