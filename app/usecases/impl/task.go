package impl

import (
	"go-todolist/app/models"
	"go-todolist/app/repositories"
	"go-todolist/app/usecases"
	"go-todolist/utils/constants"
	"gorm.io/gorm"
	"net/http"
)

type TaskUseCase struct {
	TaskRepository repositories.ITaskRepository
}

func CreateTaskUseCase(repository repositories.ITaskRepository) usecases.ITaskUseCase {
	return &TaskUseCase{TaskRepository: repository}
}

func (usecase *TaskUseCase) All(offset int, limit int) (tasks []*models.Task, code int) {
	tasks, err := usecase.TaskRepository.All(offset, limit)
	if err != nil {
		tasks = nil
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
			return
		}
		code = http.StatusInternalServerError
		return
	}

	code = constants.OK_CODE
	return
}

func (usecase *TaskUseCase) Create(task *models.Task) (createdTask *models.Task, code int) {
	createdTask, err := usecase.TaskRepository.Create(task)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return createdTask, constants.OK_CODE
}

func (usecase *TaskUseCase) Toggle(id int) (code int) {
	task, err := usecase.TaskRepository.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound
		}
		return http.StatusInternalServerError
	}

	task.Checked = !task.Checked
	task, err = usecase.TaskRepository.Update(task)
	if err != nil {
		return http.StatusInternalServerError
	}

	return constants.OK_CODE
}

func (usecase *TaskUseCase) Delete(id int) (code int) {
	_, err := usecase.TaskRepository.GetById(id)
	if err != nil {
		return http.StatusNotFound
	}

	err = usecase.TaskRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError
	}

	return constants.OK_CODE
}
