package redis

import "time"

type SessionRepository struct {
	Redishandler
}

func (repo *SessionRepository) Store(key string, value string, exp time.Duration) error {
	if err := repo.Redishandler.Set(key, value, exp); err != nil {
		return err
	}
	return nil
}

func (repo *SessionRepository) Get(key string) (string, error) {
	val, err := repo.Redishandler.Get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (repo *SessionRepository) Delete(key string) error {
	if err := repo.Redishandler.Del(key); err != nil {
		return err
	}
	return nil
}

func (repo *SessionRepository) DeleteAll() error {
	if err := repo.Redishandler.FlushAll(); err != nil {
		return err
	}
	return nil
}