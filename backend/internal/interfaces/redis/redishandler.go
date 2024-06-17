package redis

import "time"

type Redishandler interface {
	Set(string, string, time.Duration) error
	Get(string) (string, error)
	Del(string) error
	FlushAll() error
}

