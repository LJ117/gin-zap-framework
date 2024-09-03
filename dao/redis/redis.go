package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	// Get Redis Config from Viper
	redisAddr := fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
	redisPassword := viper.GetString("redis.password")
	redisDB := viper.GetInt("redis.db")
	redisPoolSize := viper.GetInt("redis.pool_size")

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
		PoolSize: redisPoolSize,
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
