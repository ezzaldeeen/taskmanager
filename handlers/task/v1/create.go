package taskv1

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/models"
)

type TaskCreatorService interface {
	Create(ido models.TaskIDO) error
}

// Create is an HTTP handler function that handles the creation of a new rest.
func (h Handler) Create(ec echo.Context) error {
	var taskDTO TaskDTO
	err := json.NewDecoder(ec.Request().Body).Decode(&taskDTO)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	taskIDO := fromDTOtoIDO(taskDTO)
	err = h.service.Create(taskIDO)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	return ec.JSON(http.StatusCreated, taskDTO)
}
