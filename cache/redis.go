package cache

import (
	"ace/model"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
)

var Client *redis.Client

func InitRedis(conf model.Redis) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		zap.L().Error("[Redis] Ping redis failure", zap.Error(err))
		os.Exit(2)
	}
	Client = rdb
}
