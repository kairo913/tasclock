package usecase

import "github.com/kairo913/tasclock/internal/domain/model"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Add(t model.Task) (id int64, err error) {
	id, err = interactor.TaskRepository.Store(t)
	return
}

func (interactor *TaskInteractor) Update(t model.Task) (err error) {
	err = interactor.TaskRepository.Update(t)
	return
}

func (interactor *TaskInteractor) Delete(t model.Task) (err error) {
	err = interactor.TaskRepository.Delete(t)
	return
}

func (interactor *TaskInteractor) FindById(id int) (task model.Task, err error) {
	task, err = interactor.TaskRepository.FindById(id)
	return
}

func (interactor *TaskInteractor) FindByUserId(userId int) (tasks model.Tasks, err error) {
	tasks, err = interactor.TaskRepository.FindByUserId(userId)
	return
}

func (interactor *TaskInteractor) FindAll() (tasks model.Tasks, err error) {
	tasks, err = interactor.TaskRepository.FindAll()
	return
}