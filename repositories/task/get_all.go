package taskrepo

import (
	"taskmanager/models"
)

func (r Repository) GetAll() ([]models.TaskIDO, error) {
	var tasks []models.TaskIDO
	conn, err := r.handlerFn()
	if err != nil {
		return nil, err
	}
	if err := conn.Table(r.table).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
