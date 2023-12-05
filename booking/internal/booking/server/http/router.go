package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vshigimoto/BookingService/internal/booking/config"
	"github.com/vshigimoto/BookingService/internal/booking/server/http/middleware"
	"github.com/vshigimoto/BookingService/internal/booking/usecase"
	"go.uber.org/zap"
)

type bookingRouter struct {
	u   usecase.BookingUC
	l   *zap.SugaredLogger
	cfg config.Config
}

func InitRouter(r *gin.Engine, u usecase.BookingUC, l *zap.SugaredLogger, cfg config.Config) {
	br := bookingRouter{
		u:   u,
		l:   l,
		cfg: cfg,
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("api/booking/v1")
	{
		v1.GET("/booking/book/:id", middleware.JWTVerify(), br.u.BookRoom())
		v1.GET("/booking/hotel/:id", middleware.JWTVerify(), br.u.GetHotelById())
		v1.GET("/booking/hotel", middleware.JWTVerify(), br.u.GetHotels())
		v1.POST("/booking/hotel/confirm", middleware.JWTVerify(), br.u.ConfirmBook())
	}
}

func InitAdminRouter(r *gin.Engine, u usecase.BookingUC, l *zap.SugaredLogger, cfg config.Config) {
	br := bookingRouter{
		u:   u,
		l:   l,
		cfg: cfg,
	}
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("api/admin/v1")
	{
		v1.POST("/booking/hotel", middleware.AdminVerify(), br.u.CreateHotel())
		v1.PUT("/booking/hotel/:id", middleware.AdminVerify(), br.u.UpdateHotel())
		v1.DELETE("/booking/hotel/:id", middleware.AdminVerify(), br.u.DeleteHotel())

	}
}
