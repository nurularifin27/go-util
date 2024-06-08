package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	Logger = logger
	defer logger.Sync()
}
