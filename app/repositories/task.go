package repositories

import "go-todolist/app/models"

type ITaskRepository interface {
	All(offset int, limit int) (tasks []*models.Task, err error)
	GetById(id uint) (task *models.Task, err error)
	Create(task *models.Task) (createdTask *models.Task, err error)
	Update(task *models.Task) (updatedTask *models.Task, err error)
	Delete(id uint) (err error)
}
