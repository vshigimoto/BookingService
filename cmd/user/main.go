package main

import (
	_ "github.com/lib/pq" //
	"github.com/vshigimoto/BookingService/internal/user/applicator"
	"github.com/vshigimoto/BookingService/internal/user/config"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	l := logger.Sugar()
	l = l.With(zap.String("app", "user-service"))
	cfg, err := config.LoadConfig("config/user")
	if err != nil {
		l.Fatalf("Failed to load config: %v", err)
	}
	app := applicator.New(cfg, l)
	app.Run()
}
