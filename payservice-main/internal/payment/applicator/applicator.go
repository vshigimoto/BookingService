package applicator

import (
	"go.uber.org/zap"

	"payservice/internal/payment/config"
	"payservice/internal/payment/database"
	"payservice/internal/payment/payment"
	"payservice/internal/payment/repository"
	"payservice/internal/payment/server/http"
)

func Run(cfg *config.Config, l *zap.SugaredLogger) {
	mainDB := database.Connect(cfg.Database.Main)
	defer mainDB.Close()

	replicaDB := database.Connect(cfg.Database.Replica)
	defer replicaDB.Close()

	queryBuilder := repository.NewSQLQueryBuilder()

	repo := repository.NewRepository(mainDB, replicaDB, queryBuilder)

	paymentService := payment.NewService(*repo, cfg.Auth.JwtSecretKey, cfg.Auth.PasswordSecretKey, l)

	endpointHandler := http.NewEndpointHandler(paymentService, l)
	http.Start(endpointHandler)
}
