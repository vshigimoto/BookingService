package applicator

import (
	"booking/internal/auth/auth"
	"booking/internal/auth/config"
	"booking/internal/auth/database"
	"booking/internal/auth/repository"
	"booking/internal/auth/server/http"
	"booking/internal/auth/transport"
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
		l.Panicf("Error to connect DB '%s':%v", cfg.Database.Main.Host, err)
	}
	replicaDB, err := database.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("Error to connect DB '%s':%v", cfg.Database.Replica.Host, err)
	}
	rep := repository.NewRepository(mainDB, replicaDB)
	userTransport := transport.NewTransport(cfg.Transport.UserTransport)
	if err != nil {
		l.Panicf("failed NewProducer err: %v", err)
	}

	authService := auth.NewAuthService(rep, cfg, userTransport)
	endpointHandler := http.NewEndpointHandler(*authService)
	http.InitRouter(rep, r, endpointHandler)
	port := fmt.Sprintf(":%d", cfg.HttpServer.Port)
	if err := r.Run(port); err != nil {
		l.Panicf("Error to run server %v", err)
	}
}
