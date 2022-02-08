package repositories

import "go-todolist/app/models"

type ITaskRepository interface {
	All() (err error)
	Create(task *models.Task) (createdTask *models.Task, err error)
	Toggle(id uint) (err error)
	Delete(id uint) (err error)
}
