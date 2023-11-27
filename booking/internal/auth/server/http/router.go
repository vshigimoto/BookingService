package http

import (
	"booking/internal/auth/repository"
	"github.com/gin-gonic/gin"
)

func InitRouter(rep *repository.Repo, r *gin.Engine, eh *EndpointHandler) {
	v1 := r.Group("api/auth/v1")
	{
		v1.POST("/user/login", eh.Login())
	}
}
