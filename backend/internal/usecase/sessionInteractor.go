package usecase

import "time"

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) Set(key string, value string, exp time.Duration) (err error) {
	err = interactor.SessionRepository.Set(key, value, exp)
	return
}

func (interactor *SessionInteractor) Get(key string) (value string, err error) {
	value, err = interactor.SessionRepository.Get(key)
	return
}

func (interactor *SessionInteractor) Del(key string) (err error) {
	err = interactor.SessionRepository.Del(key)
	return
}

func (interactor *SessionInteractor) FlushAll() (err error) {
	err = interactor.SessionRepository.FlushAll()
	return
}