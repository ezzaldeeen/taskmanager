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
	var taskReq TaskRequestModel
	err := json.NewDecoder(ec.Request().Body).Decode(&taskReq)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	taskIDO := convertToTaskIDO(taskReq)
	err = h.service.Create(taskIDO)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	return ec.JSON(http.StatusCreated, taskReq)
}
