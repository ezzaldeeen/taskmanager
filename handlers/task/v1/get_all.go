package taskv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/models"
)

type TasksGetterService interface {
	GetAll() ([]models.TaskIDO, error)
}

// GetAll is an HTTP handler function that retrieves all tasks.
func (h Handler) GetAll(c echo.Context) error {
	taskIDOs, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	tasksRes := convertTasksToResponseModels(taskIDOs)
	return c.JSON(http.StatusOK, tasksRes)
}
