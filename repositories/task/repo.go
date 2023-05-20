package taskrepo

import (
	"gorm.io/gorm"
	"taskmanager/database"
)

// tableName DB table name
const (
	tableName = "tasks"
)

// Repository represents a repository responsible
// for handling the DB accessing logic.
type Repository struct {
	table     string
	handlerFn func() (*gorm.DB, error)
}

// NewRepository creates a new instance of the Repo.
func NewRepository(driver *database.DBDriver) Repository {
	return Repository{
		table:     tableName,
		handlerFn: driver.GetConnection,
	}
}
