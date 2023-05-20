package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"taskmanager/config"
	v1 "taskmanager/controllers/rest/v1"
	taskv1 "taskmanager/controllers/task/v1"
	"taskmanager/database"
	"taskmanager/middlewares"
	taskrepo "taskmanager/repositories/task"
	tasksvc "taskmanager/services/task"
)

func main() {
	e := echo.New()
	// Load app's configuration
	cfg := config.Load()
	// Setup DB driver
	driver := database.NewDBDriver(
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
		cfg.DB.Port)

	// register pre-requests middlewares
	e.Pre(middlewares.TrimTrailingSlashes)

	// Instantiate route group
	api := e.Group("/api")

	// Initialize Repositories
	taskRepo := taskrepo.NewRepo(driver)

	// Initialize Services
	taskSvc := tasksvc.NewService(taskRepo)

	// API group registration
	taskv1.NewController(api, taskSvc).Register()
	v1.NewController(api, driver).Register()

	// Starting service based on the given configuration
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Service.Port)))
}
