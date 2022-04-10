package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/burkaydurdu/wish/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const version = "0.0.2"

func createServer() *gin.Engine {
	return gin.Default()
}

func startHttpServer(lifecycle fx.Lifecycle, g *gin.Engine, config *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			g.GET("/health", healthCheck)
			return g.Run(fmt.Sprintf(":%s", config.Port))
		},
	})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "OK", "version": version})
}

func createLogger() *zap.Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}

	zap.ReplaceGlobals(logger)
	return logger
}
