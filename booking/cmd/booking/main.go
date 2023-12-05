package main

import (
	_ "github.com/lib/pq" //
	"github.com/vshigimoto/BookingService/internal/booking/applicator"
	"github.com/vshigimoto/BookingService/internal/booking/config"
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
