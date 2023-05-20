package taskrepo

import (
	"gorm.io/gorm"
	"taskmanager/database"
)

const (
	tableName = "tasks"
)

type Repo struct {
	table     string
	handlerFn func() (*gorm.DB, error)
}

func NewRepo(driver *database.DBDriver) Repo {
	return Repo{
		table:     tableName,
		handlerFn: driver.GetConnection,
	}
}
