package usecase

import (
	"booking/internal/user/entity"
	"booking/internal/user/repository"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserUC struct {
	l *zap.SugaredLogger
	r *repository.Repo
}

func NewUserUC(l *zap.SugaredLogger, r *repository.Repo) *UserUC {
	return &UserUC{
		l: l,
		r: r,
	}
}

func (u *UserUC) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			u.l.Warnf("cannot parse user with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		id, err := u.r.CreateUser(context.TODO(), &user)
		if err != nil {
			u.l.Warnf("cannot create user with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot create user"})
			return
		}

		ctx.JSON(http.StatusOK, id)
	}
}

func (u *UserUC) GetByLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("login")
		if login == "" {
			u.l.Warnf("cannot get by login, login is empty")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "login must not be empty"})
			return
		}
		user, err := u.r.GetByLogin(context.TODO(), login)
		if err != nil {
			u.l.Warnf("cannot get from db with error:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot get from db"})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (u *UserUC) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("login")
		if login == "" {
			u.l.Warnf("cannot update user, login is empty")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "login must not be empty"})
			return
		}
		var user entity.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			u.l.Warnf("cannot unmarshall user with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "cannot unmarshall user"})
			return
		}
		err := u.r.UpdateUser(context.TODO(), login, &user)
		if err != nil {
			u.l.Warnf("cannot update user with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cannot update user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (u *UserUC) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		login := ctx.Param("login")
		if login == "" {
			u.l.Warnf("cannot update user, login is empty")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "login must not be empty"})
			return
		}

		err := u.r.DeleteUser(context.TODO(), login)
		if err != nil {
			u.l.Warnf("cannot delete user with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (u *UserUC) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sortBy := ctx.Query("sortBy")
		sortKey := ctx.Query("sortKey")
		if sortBy != "ASC" && sortBy != "DESC" && sortBy != "" {
			u.l.Warnf("cannot execute getUsers because sortBy is invalid")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "sortBy is invalid"})
			return
		}
		if sortKey == "" {
			sortKey = "id"
		}
		users, err := u.r.GetUsers(context.TODO(), sortKey, sortBy)
		if err != nil {
			u.l.Warnf("cannot get users with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get user"})
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}
