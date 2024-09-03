package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

/*
  Go Web Dev Genera Framework Template
*/

func main() {
	// 1. load config
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err: %v\n", err)
		return
	}
	// 2. init log
	if err := logger.Init(settings.SysCfg.LogConfig); err != nil {
		fmt.Printf("Init logger failed, err: %v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())

	// 3. init DB(MySQL,PG) Connect
	if err := mysql.Init(settings.SysCfg.MySQLConfig); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close() //free to connect

	// 4. init Redis Connect
	if err := redis.Init(settings.SysCfg.RedisConfig); err != nil {
		fmt.Printf("Init redis failed, err: %v\n", err)
		return
	}
	defer redis.Close() //free to connect

	// 5. register Routers
	r := routes.Setup(settings.SysCfg.AppConfig)

	// 6. start serving (graceful terminate)
	svcPort := viper.GetInt("app.port")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", svcPort),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("ListenAndServe", zap.Error(err))
		}
	}()

	//graceful leave
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
		return
	}
	zap.L().Info("Server exiting")
}
