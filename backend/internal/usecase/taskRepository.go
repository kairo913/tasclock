package usecase

import "github.com/kairo913/tasclock/internal/domain/model"

type TaskRepository interface {
	Store(model.Task) (int64, error)
	Update(model.Task) error
	Delete(model.Task) error
	FindById(int) (model.Task, error)
	FindByUserId(int) (model.Tasks, error)
	FindAll() (model.Tasks, error)
}