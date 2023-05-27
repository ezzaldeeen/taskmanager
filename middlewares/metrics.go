package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "api",
			Name:      "requests_total",
			Help:      "Total number of requests",
		},
		[]string{"method", "path", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
}

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		requestCount.WithLabelValues(
			req.Method,
			req.URL.Path,
			strconv.Itoa(res.Status),
		).Inc()
		return next(c)
	}
}
