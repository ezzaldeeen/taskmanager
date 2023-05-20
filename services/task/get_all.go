package tasksvc

import (
	"taskmanager/models"
)

// TasksGetterRepo defines the interface for a tasks getter repository.
type TasksGetterRepo interface {
	GetAll() ([]models.IDO, error)
}

// GetAll retrieve all tasks
func (s Service) GetAll() ([]models.IDO, error) {
	idos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return idos, nil
}
