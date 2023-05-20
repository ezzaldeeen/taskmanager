package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBDriver struct {
	user     string
	name     string
	pwd      string
	port     string
	sslMode  string
	timeZone string
}

func NewDBDriver(user, name, pwd, port string) *DBDriver {
	return &DBDriver{
		user: user,
		name: name,
		pwd:  pwd,
		port: port,
	}
}

// todo use sync.Once
func (d DBDriver) GetConnection() (*gorm.DB, error) {
	// setting database configuration
	dns := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s",
		d.user, d.pwd, d.name, d.port)
	cfg := postgres.Config{
		DriverName:           "",
		DSN:                  dns,
		PreferSimpleProtocol: true,  // todo: read about
		WithoutReturning:     false, // todo: read about
		Conn:                 nil,   // todo: read about connection pool, read about TCP, TCP overhead, read about db connections how it works
	}
	// getting new db connection
	db, err := gorm.Open(postgres.New(cfg))
	if err != nil {
		return nil, err
	}
	return db, nil
}
