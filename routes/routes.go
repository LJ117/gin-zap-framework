package routes

import (
	"fmt"
	"net/http"
	"web_app/logger"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Setup(cfg *settings.AppConfig) *gin.Engine {

	fmt.Printf("The current gin mode is:%s\n", cfg.GinModel)
	zap.L().Info("The current gin mode is: " + cfg.GinModel)

	gin.SetMode(cfg.GinModel)
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(cfg.PrintStackInfo))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/versions", func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Version)
	})
	return r
}
