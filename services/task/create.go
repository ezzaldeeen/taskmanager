package tasksvc

import (
	"taskmanager/models"
)

// TaskCreatorRepo defines the interface for a task creator repository.
type TaskCreatorRepo interface {
	Create(task models.TaskIDO) error
}

// Create creates a new task based on the provided IDO.
func (s Service) Create(task models.TaskIDO) error {
	err := s.repo.Create(task)
	if err != nil {
		return err
	}
	return nil
}
