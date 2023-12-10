package applicator

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vshigimoto/BookingService/internal/auth/auth"
	"github.com/vshigimoto/BookingService/internal/auth/config"
	"github.com/vshigimoto/BookingService/internal/auth/database"
	"github.com/vshigimoto/BookingService/internal/auth/repository"
	"github.com/vshigimoto/BookingService/internal/auth/server/http"
	"github.com/vshigimoto/BookingService/internal/auth/transport"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Applicator struct {
	logger *zap.SugaredLogger
	config config.Config
}

func New(cfg config.Config, logger *zap.SugaredLogger) *Applicator {
	return &Applicator{
		config: cfg,
		logger: logger,
	}
}

func (a *Applicator) Run() {
	r := gin.Default()
	cfg := a.config
	l := a.logger
	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDB, err := database.New(cfg.Database.Main)
	if err != nil {
		l.Panicf("Error to connect DB '%s':%v", cfg.Database.Main.Host, err)
	}
	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed close mainDB err: %v", err)
		}
		l.Info("mainDB closed")
	}()

	replicaDB, err := database.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("Error to connect DB '%s':%v", cfg.Database.Replica.Host, err)
	}
	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed close replicaDB err: %v", err)
		}
		l.Info("replicaDB closed")
	}()

	rep := repository.New(mainDB, replicaDB)

	userTransport := transport.New(cfg.Transport.UserTransport)
	if err != nil {
		l.Panicf("failed NewProducer err: %v", err)
	}

	authService := auth.New(rep, cfg, userTransport)
	endpointHandler := http.New(*authService)

	http.InitRouter(r, endpointHandler)
	port := fmt.Sprintf(":%d", cfg.HttpServer.Port)
	if err := r.Run(port); err != nil {
		l.Panicf("Error to run server %v", err)
	}

	a.gracefulShutdown(cancel)
}

func (a *Applicator) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
