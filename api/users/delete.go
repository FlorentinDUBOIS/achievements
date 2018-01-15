package users

import (
	"net/http"

	usersSrv "github.com/FlorentinDUBOIS/achievements/services/users"
	"github.com/labstack/echo"
)

// Delete an user
func Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := usersSrv.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
