package config

import (
	"github.com/nurularifin27/go-util/pkg/logger"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func LoadConfig() {
	logger.InitLogger()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("error load config ", zap.Error(err))
	}
}
