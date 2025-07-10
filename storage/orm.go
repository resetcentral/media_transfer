package storage

import (
	"github.com/resetcentral/media_transfer/models"
	"gorm.io/gorm"
)

type GormStorage struct {
	db *gorm.DB
}

func NewGormStorage() (GormStorage, error) {
	var storage GormStorage
	return storage, nil
}

func (s GormStorage) CreateTask(task *models.Task) (int, error) {
	return 0, nil
}

func (s GormStorage) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := s.db.Find(&tasks)
	return tasks, result.Error
}

func (s GormStorage) GetTask(task models.Task) (models.Task, error) {
	return task, nil
}

func (s GormStorage) GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	return task, nil
}

func (s GormStorage) UpdateTask(task models.Task) error {
	return nil
}

func (s GormStorage) DeleteTaskByID(id int) error {
	return nil
}
