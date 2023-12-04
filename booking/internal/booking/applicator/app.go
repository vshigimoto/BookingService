package applicator

import (
	"booking/internal/booking/config"
	"booking/internal/booking/database"
	"booking/internal/booking/kafka"
	"booking/internal/booking/repository"
	"booking/internal/booking/server/consumer"
	"booking/internal/booking/server/http"
	"booking/internal/booking/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
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
	bookingVerificationProducer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		l.Panicf("failed NewProducer err: %v", err)
	}
	bookingVerificationConsumerCallback := consumer.NewBookingVerificationCallback(l)

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
	bookingUC := usecase.NewBookingUC(l, rep, bookingVerificationProducer)
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
}
