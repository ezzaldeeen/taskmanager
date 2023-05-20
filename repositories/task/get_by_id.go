package taskrepo

import (
	"taskmanager/models"
)

// GetByID ...
func (r Repository) GetByID(id string) (models.TaskIDO, error) {
	var task models.TaskIDO
	conn, err := r.handlerFn()
	if err != nil {
		return models.TaskIDO{}, err
	}
	err = conn.Table(r.table).Where("id = ?", id).First(&task).Error
	if err != nil {
		return models.TaskIDO{}, err
	}
	return task, nil
}
