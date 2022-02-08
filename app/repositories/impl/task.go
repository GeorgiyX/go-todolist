package impl

import (
	"go-todolist/app/models"
	"go-todolist/app/repositories"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func CreateTaskRepository(db *gorm.DB) repositories.ITaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) All() (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TaskRepository) Create(task *models.Task) (createdTask *models.Task, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TaskRepository) Toggle(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *TaskRepository) Delete(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}
