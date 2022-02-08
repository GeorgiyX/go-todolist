package impl

import (
	"go-todolist/app/models"
	"go-todolist/app/repositories"
	"go-todolist/app/usecases"
)

type TaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func CreateTaskUseCase(repository repositories.ITaskRepository) usecases.ITaskUseCase {
	return &TaskUseCase{TaskRepository: repository}
}

func (usecase *TaskUseCase) All() (err error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *TaskUseCase) Create(task *models.Task) (createdTask *models.Task, err error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *TaskUseCase) Toggle(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *TaskUseCase) Delete(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}


