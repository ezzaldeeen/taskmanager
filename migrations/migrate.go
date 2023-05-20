package main

import (
	"log"
	"taskmanager/config"
	"taskmanager/database"
	"taskmanager/migrations/internal"
)

func main() {
	// load app's configuration
	cfg := config.Load()
	driver := database.NewDBDriver(
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
		cfg.DB.Port)

	conn, err := driver.GetConnection()
	if err != nil {
		log.Fatal("unable to connect with DB:", err)
	}
	// perform schema migration
	err = conn.AutoMigrate(&internal.Task{})
	if err != nil {
		log.Fatal("unable to migrate with DB:", err)
	}
}
