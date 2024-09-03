package routes

import (
	"net/http"
	"web_app/logger"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	ginModel := viper.GetString("app.gin_model")
	gin.SetMode(ginModel)
	r := gin.New()

	isPrintStackInfo := viper.GetBool("app.print_stack_info")

	r.Use(logger.GinLogger(), logger.GinRecovery(isPrintStackInfo))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
