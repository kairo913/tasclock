package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/interfaces/database"
	"github.com/kairo913/tasclock/internal/usecase"
)

type TaskController struct {
	TaskInteractor    usecase.TaskInteractor
}

func NewTaskController(sqlhandler database.Sqlhandler) *TaskController {
	return &TaskController{
		TaskInteractor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				Sqlhandler: sqlhandler,
			},
		},
	}
}

func (controller *TaskController) Create(c *gin.Context) {

}
