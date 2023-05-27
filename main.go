package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	dbDriver := database.NewDBDriver(
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
		cfg.DB.Port)

	// Register Middlewares
	e.Pre(middlewares.TrimTrailingSlashes)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middlewares.PrometheusMiddleware)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Instantiate route group
	api := e.Group("/api")

	// Initialize Repositories
	taskRepo := taskrepo.NewRepository(dbDriver)

	// Initialize Services
	taskSvc := tasksvc.NewService(taskRepo)

	// API group registration
	health.NewHandler(api).Register()
	taskv1.NewHandler(api, taskSvc).Register()

	// Starting service based on the given configuration
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Service.Port)))
}
