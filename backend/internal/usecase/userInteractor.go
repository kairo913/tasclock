package usecase

import "github.com/kairo913/tasclock/internal/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (id int64, err error) {
	id, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u model.User) (err error) {
	err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) Delete(u model.User) (err error) {
	err = interactor.UserRepository.Delete(u)
	return
}

func (interactor *UserInteractor) FindById(id int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) FindByEmail(email string) (user model.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)
	return
}

func (interactor *UserInteractor) FindAll() (users model.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

