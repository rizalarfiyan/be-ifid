package database

import (
	"be-ifid/config"
	"be-ifid/utils"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisConn *redis.Client

func RedisInit() {
	utils.Info("Connect redis server...")
	conf := config.Get()
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
		Username:    conf.Redis.User,
		Password:    conf.Redis.Password,
		DB:          0,
		DialTimeout: conf.Redis.DialTimeout,
	})

	redisConn = new(redis.Client)
	redisConn = rdb

	utils.Success("Redis connected")
}

func RedisGet() *redis.Client {
	return redisConn
}

func RedisIsConnected() bool {
	_, err := redisConn.Ping(context.Background()).Result()
	if err != nil {
		utils.SafeError("Redis fails health check")
		return false
	}
	return true
}
