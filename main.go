package main

import (
	"TodoList/config"
	"TodoList/controllers/task"
	"TodoList/middlewares"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// load app's configuration
	cfg := config.Load()
	// register pre-requests middlewares
	e.Pre(middlewares.TrimTrailingSlashes)
	// instantiate route group
	api := e.Group("/api")
	// register route groups
	task.RegisterRouteGrp(api)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Service.Port)))
}
