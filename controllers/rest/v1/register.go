package v1

import (
	"github.com/labstack/echo/v4"
	"taskmanager/controllers"
	"taskmanager/database"
	"taskmanager/repositories/task"
	"taskmanager/services/task"
)

// todo: split this file into three files
type Control controllers.Controller

// NewController creates a new instance of the Controller.
func NewController(g *echo.Group, driver *database.DBDriver) controllers.Registrar {
	repo := taskrepo.NewRepo(driver)
	svc := tasksvc.NewService(repo)
	return &Control{
		G:       g,
		Service: svc,
	}
}

// Register registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/tasks" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: Retrieve all tasks
// - GET: Retrieve a rest by its ID todo (Not implemented)
// - POST: Create a new rest
// - PUT: Update a rest by its ID todo (Not implemented)
// - DELETE: Delete a rest by its ID todo (Not implemented)
func (c Control) Register() {
	grp := c.G.Group("/v1/tasks")
	grp.GET("", c.getAll)
	grp.GET("/:id", c.getById)
	grp.POST("", c.create)
	grp.PUT("/:id", nil)
	grp.DELETE("/:id", nil)
}
