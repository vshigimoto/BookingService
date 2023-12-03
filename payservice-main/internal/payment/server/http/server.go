package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(eh *EndpointHandler) {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v2 := r.Group("v2")
	{
		v2.GET("/admin/card/payments", eh.GetAdminCardPayments)
		v2.DELETE("/admin/card/payment/:id", eh.DeleteAdminCardPayment)

		v2.Use(eh.TokenAuthMiddleware())
		{
			v2.Use(eh.MetricsHandler())
			v2.GET("/cards", eh.GetCards)
			v2.POST("/card", eh.CreateCard)
			v2.PUT("/card/:id", eh.UpdateCard)
			v2.DELETE("/card/:id", eh.DeleteCard)
			v2.POST("/card/payment", eh.CreateCardPayment)
			v2.GET("/card/payment/:id", eh.GetCardPayment)
			v2.GET("/card/payments", eh.GetCardPayments)
			v2.GET("/accounts", eh.GetAccounts)
			v2.GET("/account/:id", eh.GetAccount)
			v2.POST("/account", eh.CreateAccount)
			v2.DELETE("/account/:id", eh.DeleteAccount)
			v2.POST("/account/payment", eh.CreateAccountPayment)
			v2.GET("/account/payment/:id", eh.GetAccountPayment)
			v2.GET("/account/payments", eh.GetAccountPayments)
		}
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//nolint:all
	r.Run(":8082")
}
