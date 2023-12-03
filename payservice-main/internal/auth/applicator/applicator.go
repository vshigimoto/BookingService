package applicator

import (
	"go.uber.org/zap"

	"payservice/internal/auth/auth"
	"payservice/internal/auth/config"
	"payservice/internal/auth/controller/consumer"
	"payservice/internal/auth/database"
	"payservice/internal/auth/repository"
	"payservice/internal/auth/server/http"
	"payservice/internal/auth/transport"
	"payservice/internal/kafka"
)

func Run(cfg *config.Config, l *zap.SugaredLogger) {
	mainDB := database.Connect(cfg.Database.Main)
	defer mainDB.Close()

	replicaDB := database.Connect(cfg.Database.Replica)
	defer replicaDB.Close()

	repo := repository.NewRepository(mainDB, replicaDB)
	userGrpcTransport := transport.NewUserGrpcTransport(cfg.Transport.UserGrpc)

	//fmt.Println(cfg.Kafka)
	userVerificationProducer, err := kafka.NewProducer(cfg.Kafka)

	if err != nil {
		l.Panicf("failed NewProducer err: %v", err)
	}

	userVerificationConsumerCallback := consumer.NewUserVerificationCallback()

	userVerificationConsumer, err := kafka.NewConsumer(cfg.Kafka, userVerificationConsumerCallback)
	if err != nil {
		l.Panicf("failed NewConsumer err: %v", err)
	}

	go userVerificationConsumer.Start()

	authService := auth.NewService(*repo, cfg.Auth, userGrpcTransport, userVerificationProducer, l)

	endpointHandler := http.NewEndpointHandler(authService, l)
	http.Start(endpointHandler)
}
