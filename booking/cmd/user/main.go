package main

import (
	"booking/internal/user/applicator"
	"booking/internal/user/config"
	_ "github.com/lib/pq" //
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()
	l = l.With(zap.String("app", "user-service"))
	cfg, err := config.LoadConfig("config/user")
	if err != nil {
		l.Fatalf("Failed to load config: %v", err)
	}
	app := applicator.NewApplicator(cfg, l)
	app.Run()
}
