package main

import (
	"booking/internal/booking/applicator"
	"booking/internal/booking/config"
	_ "github.com/lib/pq" //
	"go.uber.org/zap"
)

// @title Booking service
// @version 1.0
// @description API server for Booking service

// @host localhost:9234

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	l := logger.Sugar()
	l = l.With(zap.String("app", "booking-service"))
	cfg, err := config.LoadConfig("config/booking")
	if err != nil {
		l.Fatalf("Failed to load config: %v", err)
	}
	app := applicator.NewApplicator(cfg, l)
	app.Run()
}
