package tasksvc

import (
	"taskmanager/models"
)

// TaskCreatorRepo defines the interface for a task creator repository.
type TaskCreatorRepo interface {
	Create(ido *models.IDO) error
}

// Create creates a new task based on the provided IDO.
func (s Service) Create(ido *models.IDO) error {
	err := s.repo.Create(ido)
	if err != nil {
		return err
	}
	return nil
}
