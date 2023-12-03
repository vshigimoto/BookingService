package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	_ "payservice/internal/auth/server/http/docs"

	"payservice/internal/auth/applicator"
	"payservice/internal/auth/config"
)

// @title Auth Service API
// @version 1.0
// @description Auth service API in Go using Gin Framework
// @host localhost:8081
// @BasePath /auth
func main() {
	logger, _ := zap.NewProduction()
	//nolint:all
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "payservice-auth"))

	cfg, err := loadCfg("config/auth")
	if err != nil {
		l.Fatalf("Load to load config: %v", err)
	}

	applicator.Run(&cfg, l)
}

func loadCfg(path string) (cfg config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return cfg, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	return cfg, nil
}
