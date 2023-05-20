package main

import (
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"taskmanager/config"
	"taskmanager/database"
	"taskmanager/handlers/health"
	"taskmanager/handlers/task/v1"
	"taskmanager/middlewares"
	"taskmanager/repositories/task"
	"taskmanager/services/task"
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

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	// Register Middlewares
	e.Pre(middlewares.TrimTrailingSlashes)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Instantiate route group
	api := e.Group("/api")

	// Initialize Repositories
	taskRepo := taskrepo.NewRepository(driver)

	// Initialize Services
	taskSvc := tasksvc.NewService(taskRepo)

	// API group registration
	health.NewHandler(api).Register()
	taskv1.NewHandler(api, taskSvc).Register()

	// Starting service based on the given configuration
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Service.Port)))
}
