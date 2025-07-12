package orm

import "github.com/resetcentral/media_transfer/models"

func (s GormStorage) CreateTask(task *models.Task) error {
	result := s.db.Create(task)

	return result.Error
}

func (s GormStorage) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := s.db.Find(&tasks)

	return tasks, result.Error
}

func (s GormStorage) GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	result := s.db.Find(&task, id)

	return task, result.Error
}

func (s GormStorage) UpdateTask(task models.Task) error {
	result := s.db.Save(task)

	return result.Error
}

func (s GormStorage) DeleteTaskByID(id int) error {
	result := s.db.Delete(&models.Task{}, id)

	return result.Error
}
