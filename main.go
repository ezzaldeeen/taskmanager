package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"taskmanager/config"
	"taskmanager/controllers/health"
	"taskmanager/controllers/task"
	"taskmanager/database"
	"taskmanager/middlewares"
)

func main() {
	e := echo.New()
	// load app's configuration
	cfg := config.Load()
	// setup db driver
	driver := database.NewDBDriver(
		cfg.DB.User,
		cfg.DB.Name,
		cfg.DB.Password,
		cfg.DB.Port)
	// register pre-requests middlewares
	e.Pre(middlewares.TrimTrailingSlashes)
	// instantiate route group
	api := e.Group("/api")
	// register route groups
	health.NewController(api).Register()
	task.NewController(api, driver).Register()
	// starting service based on the given configuration
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Service.Port)))
}
