package http

import (
	"booking/internal/user/config"
	"booking/internal/user/server/http/middleware"
	"booking/internal/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type userRouter struct {
	u   usecase.UserUC
	l   *zap.SugaredLogger
	cfg config.Config
}

func UserRouter(r *gin.Engine, u usecase.UserUC, l *zap.SugaredLogger, cfg config.Config) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	ur := userRouter{
		u:   u,
		l:   l,
		cfg: cfg,
	}
	v1 := r.Group("api/user/v1")
	{
		v1.GET("/user/:login", ur.u.GetByLogin())
	}

}

func AdminRouter(r *gin.Engine, u usecase.UserUC, l *zap.SugaredLogger, cfg config.Config) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := r.Group("api/admin/v1")
	{
		ur := userRouter{
			u:   u,
			l:   l,
			cfg: cfg,
		}
		v1.GET("/user", middleware.AdminVerify(), ur.u.GetUsers())
		v1.POST("/user", middleware.AdminVerify(), ur.u.CreateUser())
		v1.PUT("/user/:login", middleware.AdminVerify(), ur.u.UpdateUser())
		v1.DELETE("/user/:login", middleware.AdminVerify(), ur.u.DeleteUser())
	}
}
