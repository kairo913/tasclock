package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock/internal/domain/model"
	"github.com/kairo913/tasclock/internal/interfaces/database"
	"github.com/kairo913/tasclock/internal/usecase"
)

type TaskController struct {
	TaskInteractor usecase.TaskInteractor
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
	var body struct {
		Title       string     `json:"title" binding:"required"`
		Description *string    `json:"description"`
		Deadline    *time.Time `json:"deadline"`
		Reward      *float64   `json:"reward"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uId, _ := c.Get("userId")

	userId, err := strconv.Atoi(uId.(string))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	task := model.Task{
		UserId:      userId,
		Title:       body.Title,
		Description: body.Description,
		Deadline:    body.Deadline,
		Reward:      body.Reward,
	}

	id, err := controller.TaskInteractor.TaskRepository.Store(task)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	task, err = controller.TaskInteractor.TaskRepository.FindById(int(id))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	task.UserId = -1

	c.JSON(http.StatusCreated, task)
}

func (controller *TaskController) Get(c *gin.Context) {
	uId, _ := c.Get("userId")

	userId, err := strconv.Atoi(uId.(string))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	tasks, err := controller.TaskInteractor.TaskRepository.FindByUserId(userId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	for i := range tasks {
		tasks[i].UserId = -1
	}

	c.JSON(http.StatusOK, tasks)
}

func (controller *TaskController) Update(c *gin.Context) {
	var body struct {
		Id          int        `json:"id" binding:"required"`
		Title       string     `json:"title" binding:"required"`
		Description *string    `json:"description"`
		Deadline    *time.Time `json:"deadline"`
		Reward      *float64   `json:"reward"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uId, _ := c.Get("userId")

	userId, err := strconv.Atoi(uId.(string))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	task, err := controller.TaskInteractor.TaskRepository.FindById(body.Id)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	if task.UserId != userId {
		c.Status(http.StatusForbidden)
		return
	}

	task.Title = body.Title
	task.Description = body.Description
	task.Deadline = body.Deadline
	task.Reward = body.Reward

	err = controller.TaskInteractor.TaskRepository.Update(task)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *TaskController) Delete(c *gin.Context) {
	var body struct {
		Id int `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	uId, _ := c.Get("userId")

	userId, err := strconv.Atoi(uId.(string))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	task, err := controller.TaskInteractor.TaskRepository.FindById(body.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	if task.UserId != userId {
		c.Status(http.StatusForbidden)
		return
	}

	err = controller.TaskInteractor.TaskRepository.Delete(task)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}