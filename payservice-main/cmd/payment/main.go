package main

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"payservice/internal/payment/applicator"
	"payservice/internal/payment/config"

	_ "payservice/internal/payment/server/http/docs"
)

func main() {
	logger, _ := zap.NewProduction()
	//nolint:all
	defer logger.Sync()

	l := logger.Sugar()
	l = l.With(zap.String("app", "payservice-payment"))

	cfg, err := loadCfg("config/payment")
	if err != nil {
		l.Fatalf("Load to load config: %v", err)
	}

	applicator.Run(&cfg, l)
}

func loadCfg(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	return config, nil
}
