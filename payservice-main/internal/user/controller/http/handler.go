package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"payservice/internal/user/entity"
	"payservice/internal/user/repository"
)

// GetAllUser		godoc
// @Summary			Get All users in service
// @Description		Get information about all users in system
// @Produce			application/json
// @Success			200
// @Router			/users [get]
func GetAllUsers(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := repo.GetAllUsers()
		ctx.JSON(http.StatusOK, result)
	}
}

// GetAllUser		godoc
// @Summary			Get user by id in service
// @Description		Getinformation about user in system
// @Param			request body int true "GetUser"
// @Produce			application/json
// @Success			200
// @Router			/user/:id [get]
func GetUserByID(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		result := repo.GetUserByID(id)
		ctx.JSON(http.StatusOK, result)
	}
}

// CreateUser		godoc
// @Summary			Create user in service
// @Description		Create user in service, for admin
// @Param			request body User true "CreateUser"
// @Produce			application/json
// @Success			200
// @Router			/user [post]
func CreateUser(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
			return
		}

		id, err := repo.CreateUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"id": id})
	}
}

// DeleteUser		godoc
// @Summary			Delete user
// @Description		Create user in service, for admin
// @Param			request body int true "DeleteUser"
// @Produce			application/json
// @Success			200
// @Router			/user/:id [delete]
func DeleteUser(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := repo.DeleteUser(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "User wasnt deleted!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Success": "User was deleted!"})
	}
}

// UpdateUser		godoc
// @Summary			Update user in service
// @Description		Gets id param and update user in service, for admin
// @Param			request body User true "CreateUser"
// @Produce			application/json
// @Success			200
// @Router			/user/:id [put]
func UpdateUser(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user entity.User

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
			return
		}

		err = repo.UpdateUser(id, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"succesfull": "User was updated!"})
	}
}

// GetUserByLogin		godoc
// @Summary			Get user by login in service
// @Description		Get user by login in system for auth system
// @Produce			application/json
// @Success			200
// @Router			/user/login/:login [get]
func GetUserByLogin(repo *repository.Repo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("login")
		result := repo.GetUserByLogin(login)
		ctx.JSON(http.StatusOK, result)
	}
}
