package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(eh *EndpointHandler) {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	eh.logger.Infof("Auth service is run")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := r.Group("auth")
	{
		auth.Use(eh.MetricsHandler())
		auth.POST("/login", eh.Login)
		auth.DELETE("/logout", eh.Logout)
		auth.PUT("/refresh", eh.Refresh)
		auth.POST("/register", eh.Register)
		auth.POST("/confirm", eh.ConfirmUser)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//nolint:all
	r.Run(":8081")
}
