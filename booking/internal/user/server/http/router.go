package http

import (
	"booking/internal/user/repository"
	"booking/internal/user/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(rep *repository.Repo, r *gin.Engine) {
	v1 := r.Group("api/user/v1")
	{
		v1.GET("/user/all", middleware.JWTVerify(), rep.GetUsers())
		v1.POST("/user", rep.CreateUser())
		v1.PUT("/user/:id", rep.UpdateUser())
		v1.DELETE("/user/:id", rep.DeleteUser())
		v1.GET("/user/:login", rep.GetByLogin())
	}

}
