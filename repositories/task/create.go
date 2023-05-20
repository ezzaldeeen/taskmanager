package taskrepo

import (
	"taskmanager/models"
)

func (r Repository) Create(task models.TaskIDO) error {
	conn, err := r.handlerFn()
	if err != nil {
		return err
	}
	res := conn.Table(r.table).Create(task)
	return res.Error
}
