package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"path"
)

// TrimTrailingSlashes is a middleware function that removes
// trailing slashes from the request URL. It redirects the request to
// the cleaned URL if there is a trailing slash.
// Returns nil if the request is redirected, or executes the next handler in the chain.
func TrimTrailingSlashes(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		w := c.Response().Writer
		url := r.URL.String()
		cleanPath := path.Clean(url)
		if url != cleanPath && len(url) > 1 {
			http.Redirect(w, r, cleanPath, http.StatusMovedPermanently)
			return nil
		}
		return next(c)
	}
}
