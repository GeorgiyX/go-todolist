package impl

import (
	"go-todolist/app/models"
	"go-todolist/app/repositories"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func CreateTaskRepository(db *gorm.DB) repositories.ITaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) GetById(id int) (task *models.Task, err error) {
	task = &models.Task{}
	err = repo.db.First(task, id).Error
	return
}

func (repo *TaskRepository) All(offset int, limit int) (tasks []*models.Task, err error) {
	tasks = make([]*models.Task, 0)
	err = repo.db.Limit(limit).Offset(offset).Find(&tasks).Error
	return
}

func (repo *TaskRepository) Create(task *models.Task) (createdTask *models.Task, err error) {
	err = repo.db.Create(task).Error
	createdTask = &models.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Checked:     task.Checked,
	}
	return
}

func (repo *TaskRepository) Update(task *models.Task) (updatedTask *models.Task, err error) {
	updatedTask = &models.Task{}
	err = repo.db.Model(updatedTask).Clauses(clause.Returning{}).Save(task).Error
	return
}

func (repo *TaskRepository) Delete(id int) (err error) {
	err = repo.db.Delete(&models.Task{}, id).Error
	return
}
