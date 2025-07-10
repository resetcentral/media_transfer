package storage

import "github.com/resetcentral/media_transfer/models"

type Storage interface {
	CreateTask(models.Task) (int, error)
	GetTasks() ([]models.Task, error)
	GetTask(models.Task) (models.Task, error)
	GetTaskByID(int) (models.Task, error)
	UpdateTask(models.Task) error
	DeleteTaskByID(int) error
}

var DB Storage
