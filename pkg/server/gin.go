package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nurularifin27/go-util/pkg/logger"
	"go.uber.org/zap"
)

func NewGinServer() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	return r
}

func RunServer(r *gin.Engine, port string) {
	if err := r.Run(port); err != nil {
		logger.Logger.Fatal("could not start server", zap.Error(err))
	}
}
