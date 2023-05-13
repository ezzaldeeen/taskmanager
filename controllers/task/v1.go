package task

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegisterRouteGrp registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/tasks" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: Retrieve all tasks
// - GET: Retrieve a task by its ID
// - POST: Create a new task
// - PUT: Update a task by its ID todo (Not implemented)
// - DELETE: Delete a task by its ID todo (Not implemented)
func RegisterRouteGrp(g *echo.Group) {
	taskGrp := g.Group("/v1/tasks")
	taskGrp.GET("", getAll)
	taskGrp.GET("/:id", getById)
	taskGrp.POST("", create)
	taskGrp.PUT("/:id", nil)
	taskGrp.DELETE("/:id", nil)
}

// getAll is an HTTP handler function that retrieves all tasks.
func getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, "TASKS")
}

// getById is an HTTP handler function that retrieves a task by its ID.
func getById(c echo.Context) error {
	return c.JSON(http.StatusOK, "Task By Id")
}

// create is an HTTP handler function that handles the creation of a new task.
func create(c echo.Context) error {
	//dto := task.DTO{}
	//err := json.NewDecoder(c.Request().Body).Decode(&dto)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	//
	//task3.NewRepo()
	//svc := task2.NewService()
	//err = svc.CreateTask(dto.ToIDO())
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}
	//return c.JSON(http.StatusCreated, nil)

	return nil
}
