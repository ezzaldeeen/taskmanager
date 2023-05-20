package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/controllers"
)

// Controller represents a controller responsible
// for handling requests related to health.
type Controller struct {
	g *echo.Group
}

// NewController creates a new instance of the Controller.
func NewController(g *echo.Group) controllers.Registrar {
	return &Controller{
		g: g,
	}
}

// Register registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/health" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: Retrieve health of the service
func (c Controller) Register() {
	grp := c.g.Group("/v1/health")
	grp.GET("", c.get)
}

// getAll is an HTTP handler function that the health of the service.
func (c Controller) get(ec echo.Context) error {
	// todo: implement health checking
	// todo: check DB, Redis connection status
	return ec.JSON(http.StatusOK, "OK")
}