package database

import (
	"be-ifid/config"
	"be-ifid/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisInstance struct {
	conn *redis.Client
	conf *config.Config
	ctx  context.Context
}

type RedisInstance interface {
	Ping() error
	Get(key string, dest interface{}) error
	GetString(key string) (string, error)
	Set(key string, val interface{}) error
	Setx(key string, val interface{}) error
	Setxc(key string, expire time.Duration, val interface{}) error
	HashSet(key string, val interface{}) error
	HashSetField(key string, field string, val interface{}) error
	HashGetAll(key string) (map[string]string, error)
	HashGet(key string, field string) (string, error)
	Del(key string) error
	DelKeysByPatern(patern string) error
	Close() error
}

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

func RedisConnection(ctx context.Context) RedisInstance {
	return &redisInstance{
		conn: redisConn,
		conf: config.Get(),
		ctx:  ctx,
	}
}

func RedisIsConnected() bool {
	ctx := context.Background()
	err := RedisConnection(ctx).Ping()
	if err != nil {
		utils.SafeError("Redis fails health check", err)
		return false
	}
	return true
}

func (r *redisInstance) Ping() error {
	_, err := r.conn.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisInstance) Get(key string, dest interface{}) error {
	err := utils.MustBePointer(dest, "dest")
	if err != nil {
		return err
	}
	err = r.Ping()
	if err != nil {
		return err
	}
	val, err := redisConn.Get(r.ctx, key).Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(val, dest)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisInstance) GetString(key string) (string, error) {
	err := r.Ping()
	if err != nil {
		return "", err
	}
	val, err := redisConn.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisInstance) Set(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return redisConn.Set(r.ctx, key, val, 0).Err()
}

func (r *redisInstance) Setx(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return redisConn.Set(r.ctx, key, val, r.conf.Redis.ExpiredDuration).Err()
}

func (r *redisInstance) Setxc(key string, expire time.Duration, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return redisConn.Set(r.ctx, key, val, expire).Err()
}

func (r *redisInstance) HashSet(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.HSet(r.ctx, key, val).Err()
}

func (r *redisInstance) HashSetField(key string, field string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.HSet(r.ctx, key, field, val).Err()
}

func (r *redisInstance) HashGetAll(key string) (map[string]string, error) {
	err := r.Ping()
	if err != nil {
		return nil, err
	}
	data, err := r.conn.HGetAll(r.ctx, key).Result()
	return data, err
}

func (r *redisInstance) HashGet(key string, field string) (string, error) {
	err := r.Ping()
	if err != nil {
		return "", err
	}
	data, err := r.conn.HGet(r.ctx, key, field).Result()
	return data, err
}

func (r *redisInstance) Del(key string) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.Del(r.ctx, key).Err()
}

func (r *redisInstance) DelKeysByPatern(patern string) error {
	err := r.Ping()
	if err != nil {
		return err
	}

	val, err := r.conn.Keys(r.ctx, patern).Result()
	if err != nil {
		return err
	}
	return r.conn.Del(r.ctx, val...).Err()
}

func (r *redisInstance) Close() error {
	return r.conn.Close()
}
