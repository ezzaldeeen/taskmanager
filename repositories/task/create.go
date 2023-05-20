package taskrepo

import (
	"taskmanager/models"
)

func (r Repository) Create(ido *models.IDO) error {
	conn, err := r.handlerFn()
	if err != nil {
		return err
	}
	res := conn.Table(r.table).Create(ido)
	return res.Error
}
