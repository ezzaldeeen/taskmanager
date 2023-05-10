package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegisterRouteGrp registers the route group for tasks under the provided echo.Group.
// It creates a subgroup "/v1/health" under the given group and registers
// the relevant HTTP handlers.
// The following routes are registered:
// - GET: the health status of the service
func RegisterRouteGrp(g *echo.Group) {
	taskGrp := g.Group("/v1/health")
	taskGrp.GET("", get)
}

// get is an HTTP handler function that check the health of the service.
func get(c echo.Context) error {
	// todo: implement health checking
	// todo: check DB, Redis connection status
	return c.JSON(http.StatusOK, "OK")
}
