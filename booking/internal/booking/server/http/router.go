package http

import (
	"booking/internal/booking/repository"
	"booking/internal/booking/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, rep *repository.Repo) {
	v1 := r.Group("api/booking/v1")
	{
		v1.POST("/booking/create", rep.CreateHotel())
		v1.GET("/booking/read/all", rep.GetHotel())
		v1.PUT("/booking/update/:id", rep.UpdateHotel())
		v1.DELETE("/booking/delete/:id", rep.DeleteHotel())
		v1.GET("/booking/read/:id", rep.GetHotelById())
		v1.POST("/booking/book/:id", middleware.JWTVerify(), rep.BookRoom())

	}
}
