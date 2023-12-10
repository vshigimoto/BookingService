package applicator

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vshigimoto/BookingService/internal/booking/config"
	"github.com/vshigimoto/BookingService/internal/booking/database"
	"github.com/vshigimoto/BookingService/internal/booking/kafka"
	"github.com/vshigimoto/BookingService/internal/booking/repository"
	"github.com/vshigimoto/BookingService/internal/booking/server/consumer"
	"github.com/vshigimoto/BookingService/internal/booking/server/http"
	"github.com/vshigimoto/BookingService/internal/booking/usecase"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
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

	bookingVerificationProducer, err := kafka.New(cfg.Kafka)

	if err != nil {
		l.Panicf("failed New err: %v", err)
	}

	bookingVerificationConsumerCallback := consumer.New(l)

	bookingVerificationConsumer, err := kafka.NewConsumer(l, cfg.Kafka, bookingVerificationConsumerCallback)
	if err != nil {
		l.Panicf("failed NewConsumer err: %v", err)
	}

	var wg sync.WaitGroup
	var mu sync.RWMutex
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go rep.Worker(&wg, i)
	}
	wg.Wait()
	go repository.Daemon()
	go rep.Hotels(&mu)
	go bookingVerificationConsumer.Start()
	bookingUC := usecase.New(l, rep, bookingVerificationProducer)
	http.InitRouter(r, *bookingUC, l, cfg)
	port := fmt.Sprintf(":%d", cfg.HttpServer.Port)
	go func() {
		ar := gin.Default()
		adminPort := fmt.Sprintf(":%d", cfg.HttpServer.AdminPort)
		http.InitAdminRouter(ar, *bookingUC, l, cfg)
		if err := ar.Run(adminPort); err != nil {
			l.Panicf("cannot start admin server with error %v", err)
		}
	}()
	err = r.Run(port)
	if err != nil {
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
