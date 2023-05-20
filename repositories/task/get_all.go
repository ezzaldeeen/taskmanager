package taskrepo

import (
	"taskmanager/models"
)

func (r Repository) GetAll() ([]models.IDO, error) {
	var idos []models.IDO
	conn, err := r.handlerFn()
	if err != nil {
		return nil, err
	}
	if err := conn.Table(r.table).Find(&idos).Error; err != nil {
		return nil, err
	}
	return idos, nil
}
