package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/kairo913/tasclock/internal/env"
	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	Conn *redis.Client
}

func checkRedisConnect(rdb *redis.Client, count int) error {
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		count--
		if count == 0 {
			return err
		}
		checkRedisConnect(rdb, count)
	}
	return nil
}

func NewRedishandler() *RedisHandler {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	retry, err := env.GetEnvAsIntOrFallback("CONN_RETRY_COUNT", 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = checkRedisConnect(rdb, retry)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &RedisHandler{Conn: rdb}
}

func (handler *RedisHandler) Set(key string, value string, exp time.Duration) error {
	if err := handler.Conn.Set(context.Background(), key, value, exp).Err(); err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) Get(key string) (string, error) {
	val, err := handler.Conn.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (handler *RedisHandler) Del(key string) error {
	if err := handler.Conn.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}

func (handler *RedisHandler) FlushAll() error {
	if err := handler.Conn.FlushAll(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}