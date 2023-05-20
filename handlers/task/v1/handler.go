package taskv1

import (
	"github.com/labstack/echo/v4"
	"taskmanager/handlers"
)

// TaskService wrapper of the business logic
type TaskService interface {
	TaskCreatorService
	TasksGetterService
}

// Handler represents a controller responsible
// for handling requests related to tasks.
type Handler struct {
	g       *echo.Group
	service TaskService
}

// NewHandler creates a new instance of the Controller.
func NewHandler(g *echo.Group, service TaskService) handlers.Registrar {
	return &Handler{
		g:       g,
		service: service,
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
func (h Handler) Register() {
	grp := h.g.Group("/v1/tasks")
	grp.GET("", h.GetAll)
	//grp.GET("/:id", h.getById)
	grp.POST("", h.Create)
	grp.PUT("/:id", nil)
	grp.DELETE("/:id", nil)
}
