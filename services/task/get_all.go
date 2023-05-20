package tasksvc

import (
	"taskmanager/models"
)

// TasksGetterRepo defines the interface for a tasks getter repository.
type TasksGetterRepo interface {
	GetAll() ([]models.TaskIDO, error)
}

// GetAll retrieve all tasks
func (s Service) GetAll() ([]models.TaskIDO, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
