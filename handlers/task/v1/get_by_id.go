package taskv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/models"
)

type TaskGetterByIDService interface {
	GetByID(id string) (models.TaskIDO, error)
}

// GetByID is an HTTP handler function that
// retrieves a task by its ID.
func (h Handler) GetByID(c echo.Context) error {
	taskIDO, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}
	taskRes := convertToResponseModel(taskIDO)
	return c.JSON(http.StatusOK, taskRes)
}
