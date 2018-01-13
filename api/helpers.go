package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// NotImplemented handler
func NotImplemented(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
