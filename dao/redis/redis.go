package redis

import (
	"context"
	"fmt"
	"web_app/settings"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DBIndex,
		PoolSize: cfg.PoolSize,
	})

	if _, err = rdb.Ping(context.Background()).Result(); err != nil {
		zap.L().Error("Redis connect failed", zap.Error(err))
		return
	}
	return
}
func Close() {
	_ = rdb.Close()
}
