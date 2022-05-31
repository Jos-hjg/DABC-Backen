package redis

import (
	"context"
	"dabc/config"
	"dabc/controllers"
	"github.com/go-redis/redis/v8"
)

func InitRedis() (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Host + ":" + config.C.Redis.Port,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	controllers.Redis = rdb
	return rdb, nil
}
