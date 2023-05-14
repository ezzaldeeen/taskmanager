package task

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/database"
	"taskmanager/models/task"
	"taskmanager/repositories/task"
	"taskmanager/services/task"
	"taskmanager/utils"
)

type Controller struct {
	g       *echo.Group
	service *tasksvc.Service
}

func NewController(g *echo.Group, driver *database.DBDriver) *Controller {
	repo := taskrepo.NewRepo(driver)
	svc := tasksvc.NewService(repo)
	return &Controller{
		g:       g,
		service: svc,
	}
}

// Register registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/tasks" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: Retrieve all tasks
// - GET: Retrieve a task by its ID
// - POST: Create a new task
// - PUT: Update a task by its ID todo (Not implemented)
// - DELETE: Delete a task by its ID todo (Not implemented)
func (c Controller) Register() {
	grp := c.g.Group("/v1/tasks")
	grp.GET("", c.getAll)
	grp.GET("/:id", c.getById)
	grp.POST("", c.create)
	grp.PUT("/:id", nil)
	grp.DELETE("/:id", nil)
}

// getAll is an HTTP handler function that retrieves all tasks.
func (c Controller) getAll(ec echo.Context) error {
	idos, err := c.service.GetAllTask()
	if err != nil {
		return ec.JSON(http.StatusNotFound, err)
	}
	dtos := utils.ConvertIDOStoDTOS(idos)
	return ec.JSON(http.StatusCreated, dtos)
}

// getById is an HTTP handler function that retrieves a task by its ID.
func (c Controller) getById(ec echo.Context) error {
	return ec.JSON(http.StatusOK, "Task By Id")
}

// create is an HTTP handler function that handles the creation of a new task.
func (c Controller) create(ec echo.Context) error {
	payload := new(task.DTO)
	// todo: add validation
	err := json.NewDecoder(ec.Request().Body).Decode(&payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	err = c.service.CreateTask(payload.ToIDO())
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	return ec.JSON(http.StatusCreated, payload)
}
