package taskv1

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/models"
)

type TaskCreatorService interface {
	Create(ido *models.IDO) error
}

// Create is an HTTP handler function that handles the creation of a new rest.
func (c Controller) Create(ec echo.Context) error {
	var payload Task
	// todo: add validation
	err := json.NewDecoder(ec.Request().Body).Decode(&payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	err = c.service.Create(payload.ToIDO())
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}
	return ec.JSON(http.StatusCreated, payload)
}
