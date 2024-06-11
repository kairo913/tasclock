package usecase

import "github.com/kairo913/tasclock/internal/domain/model"

type UserRepository interface {
	Store(model.User) (int64, error)
	Update(model.User) error
	Delete(model.User) error
	FindById(int) (model.User, error)
	FindByEmail(string) (model.User, error)
	FindAll() (model.Users, error)
}