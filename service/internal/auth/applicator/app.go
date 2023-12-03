package applicator

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"service/internal/auth/auth"
	"service/internal/auth/config"
	"service/internal/auth/controller/consumer"
	"service/internal/auth/controller/http"
	"service/internal/auth/database"
	"service/internal/auth/repository"
	"service/internal/auth/transport"
	"service/internal/kafka"
)

type Applicator struct {
	logger *zap.SugaredLogger
	config *config.Config
}

func NewApplicator(logger *zap.SugaredLogger, config *config.Config) *Applicator {
	return &Applicator{
		logger: logger,
		config: config,
	}
}

func (a *Applicator) Run() {
	var (
		cfg = a.config
		l   = a.logger
	)

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDB, err := database.New(cfg.Database.Main)
	if err != nil {
		l.Panicf("cannot сonnect to mainDB '%s:%d': %v", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed close MainDB err: %v", err)
		}
		l.Info("MainDB closed")
	}()

	replicaDB, err := database.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("cannot сonnect to replicaDB '%s:%d': %v", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed close replicaDB err: %v", err)
		}
		l.Info("replicaDB closed")
	}()

	userVerificationProducer, err := kafka.NewProducer(cfg.Kafka)
	if err != nil {
		l.Panicf("failed NewProducer err: %v", err)
	}

	userVerificationConsumerCallback := consumer.NewUserVerificationCallback(l)

	userVerificationConsumer, err := kafka.NewConsumer(l, cfg.Kafka, userVerificationConsumerCallback)
	if err != nil {
		l.Panicf("failed NewConsumer err: %v", err)
	}

	go userVerificationConsumer.Start()

	repo := repository.NewRepository(mainDB, replicaDB)

	userTransport := transport.NewTransport(cfg.Transport.User)
	userGrpcTransport := transport.NewUserGrpcTransport(cfg.Transport.UserGrpc)

	authService := auth.NewAuthService(repo, cfg.Auth, userTransport, userGrpcTransport, userVerificationProducer)

	endpointHandler := http.NewEndpointHandler(authService, l)

	router := http.NewRouter(l)
	httpCfg := cfg.HttpServer
	server, err := http.NewServer(httpCfg.Port, httpCfg.ShutdownTimeout, router, l, endpointHandler)
	if err != nil {
		l.Panicf("failed to create server err: %v", err)
	}

	server.Run()
	defer func() {
		if err := server.Stop(); err != nil {
			l.Panicf("failed close server err: %v", err)
		}
		l.Info("server closed")
	}()

	a.gracefulShutdown(cancel)
}

func (a *Applicator) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
