package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskmanager/handlers"
)

// Handler represents a controller responsible
// for handling requests related to health.
type Handler struct {
	g *echo.Group
}

// NewHandler creates a new instance of the Handler.
func NewHandler(g *echo.Group) handlers.Registrar {
	return &Handler{
		g: g,
	}
}

// Register registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/health" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: Retrieve health of the service
func (h Handler) Register() {
	grp := h.g.Group("/health")
	grp.GET("", h.get)
}

// get is an HTTP handler function that the health of the service.
func (h Handler) get(c echo.Context) error {
	// todo: implement health checking
	// todo: check DB, Redis connection status
	return c.JSON(http.StatusOK, "OK")
}
