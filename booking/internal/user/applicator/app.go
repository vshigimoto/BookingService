package applicator

import (
	"booking/internal/user/config"
	"booking/internal/user/database"
	"booking/internal/user/repository"
	"booking/internal/user/server/http"
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
	http.InitRouter(rep, r)
	port := fmt.Sprintf(":%d", cfg.HttpServer.Port)
	err = r.Run(port)
	if err != nil {
		l.Fatalf("Error to run server %v", err)
	}
}
