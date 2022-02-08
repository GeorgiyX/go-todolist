package usecases

import (
	"go-todolist/app/models"
)

type ITaskUseCase interface {
	All(offset int, limit int) (tasks []*models.Task, code int)
	Create(task *models.Task) (createdTask *models.Task, code int)
	Toggle(id int) (code int)
	Delete(id int) (code int)
}
