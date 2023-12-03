package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"payservice/internal/user/repository"
)

func Start(repo *repository.Repo) {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("v1")
	{
		v1.GET("/users", GetAllUsers(repo))
		v1.GET("/user/:id", GetUserByID(repo))
		v1.GET("/user/login/:login", GetUserByLogin(repo))
		v1.POST("/user", CreateUser(repo))
		v1.DELETE("/user/:id", DeleteUser(repo))
		v1.PUT("/user/:id", UpdateUser(repo))
	}

	//nolint:all
	r.Run(":8080")
}
