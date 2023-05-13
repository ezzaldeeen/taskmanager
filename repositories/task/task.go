package taskrepo

import (
	"taskmanager/database"
	"taskmanager/models/task"
)

const (
	tableName = "tasks"
)

type Repo struct {
	driver *database.DBDriver
	table  string
}

func NewRepo(driver *database.DBDriver) *Repo {
	return &Repo{
		driver: driver,
		table:  tableName,
	}
}

func (r Repo) GetAll() ([]task.IDO, error) {
	var idos []task.IDO
	conn, err := r.driver.GetConnection()
	if err != nil {
		return nil, err
	}
	if err := conn.Table(r.table).Find(&idos).Error; err != nil {
		return nil, err
	}
	return idos, nil
}

func (r Repo) Create(ido *task.IDO) error {
	conn, err := r.driver.GetConnection()
	if err != nil {
		return err
	}
	res := conn.Table(r.table).Create(ido)
	return res.Error
}
