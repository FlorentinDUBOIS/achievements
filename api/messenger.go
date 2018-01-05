package api

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/router"
	"github.com/labstack/echo"
)

// Messenger routes declaration
var Messenger = &router.Routes{
	router.Route{
		Method: http.MethodGet,
		Path:   "/messenger/",
		Handler: func(ctx echo.Context) error {
			return echo.NewHTTPError(http.StatusNotImplemented, nil)
		},
	},
}
