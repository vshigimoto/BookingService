package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter(r *gin.Engine, eh *EndpointHandler) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("api/auth/v1")
	{
		v1.POST("/user/login", eh.Login())
	}
}
