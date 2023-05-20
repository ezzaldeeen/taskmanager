package tasksvc

import "taskmanager/models"

// TaskGetterByIDRepo defines the interface for
// a task getter by id repository.
type TaskGetterByIDRepo interface {
	GetByID(id string) (models.TaskIDO, error)
}

// GetByID is a Service function that
// retrieves a task by its ID.
func (s Service) GetByID(id string) (models.TaskIDO, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return models.TaskIDO{}, err
	}
	return task, nil
}
