package redis

import (
	"context"
	"dabc/config"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis() (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.C.Redis.Host + ":" + config.C.Redis.Port,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})
	_, err := rdb.Get(ctx, "test1").Result()
	if err != nil {
		return nil, err
	}
	Redis = rdb
	return rdb, nil
}
