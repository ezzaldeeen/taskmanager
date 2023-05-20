package taskv1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/models"
	"taskmanager/utils"
)

type TasksGetterService interface {
	GetAll() ([]models.IDO, error)
}

// getAll is an HTTP handler function that retrieves all tasks.
func (c Controller) getAll(ec echo.Context) error {
	idos, err := c.service.GetAll()
	if err != nil {
		return ec.JSON(http.StatusNotFound, err)
	}
	dtos := utils.ConvertIDOStoDTOS(idos)
	return ec.JSON(http.StatusCreated, dtos)
}
