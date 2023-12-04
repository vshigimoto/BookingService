package applicator

import (
	"booking/internal/user/config"
	"booking/internal/user/database"
	"booking/internal/user/repository"
	"booking/internal/user/server/http"
	"booking/internal/user/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Applicator struct {
	logger *zap.SugaredLogger
	config config.Config
}

func NewApplicator(cfg config.Config, logger *zap.SugaredLogger) *Applicator {
	return &Applicator{
		config: cfg,
		logger: logger,
	}
}

func (a *Applicator) Run() {
	r := gin.Default()
	cfg := a.config
	l := a.logger
	mainDB, err := database.New(cfg.Database.Main)
	if err != nil {
		l.Fatalf("Error to connect DB '%s':%v", cfg.Database.Main.Host, err)
	}
	replicaDB, err := database.New(cfg.Database.Replica)
	if err != nil {
		l.Fatalf("Error to connect DB '%s':%v", cfg.Database.Replica.Host, err)
	}
	rep := repository.NewRepository(mainDB, replicaDB)

	userUC := usecase.NewUserUC(l, rep)
	http.UserRouter(r, *userUC, l, cfg)
	port := fmt.Sprintf(":%d", cfg.HttpServer.Port)
	go func() {
		ar := gin.Default()
		adminPort := fmt.Sprintf(":%d", cfg.HttpServer.AdminPort)
		l.Infof("Admin server on port %s, running", adminPort)
		adminUC := usecase.NewUserUC(l, rep)
		http.AdminRouter(ar, *adminUC, l, cfg)
		if err := ar.Run(adminPort); err != nil {
			l.Panicf("Error to run admin server: %v", err)
		}
	}()
	l.Infof("User server on port %s, running", port)
	if err := r.Run(port); err != nil {
		l.Panicf("Error to run server %v", err)
	}
}
