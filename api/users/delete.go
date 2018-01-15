package users

import (
	"net/http"

	usersSrv "github.com/FlorentinDUBOIS/achievements/services/users"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Delete an user
func Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := usersSrv.Delete(id); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot delete the user")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
