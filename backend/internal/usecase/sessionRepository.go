package usecase

import "time"

type SessionRepository interface {
	Set(key string, value string, exp time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	FlushAll() error
}