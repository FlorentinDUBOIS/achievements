package middlewares

import (
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	logrus.WithFields(map[string]interface{}{
		"date":       time.Now().Format(time.RFC3339),
		"remote":     c.RealIP(),
		"host":       req.Host,
		"uri":        req.RequestURI,
		"method":     req.Method,
		"path":       p,
		"referer":    req.Referer(),
		"user_agent": req.UserAgent(),
		"status":     res.Status,
		"latency":    stop.Sub(start).String(),
	}).Info("Handled request")

	return nil
}

// Logger middleware
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return handler(ctx, next)
		}
	}
}
