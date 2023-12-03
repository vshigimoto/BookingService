package usecase

import "github.com/gin-gonic/gin"

type UserUseCase interface {
	CreateUser() gin.HandlerFunc
	GetByLogin() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
}
