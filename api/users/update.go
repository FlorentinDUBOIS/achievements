package users

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/models"
	usersSrv "github.com/FlorentinDUBOIS/achievements/services/users"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Update an user
func Update(ctx echo.Context) error {
	id := ctx.Param("id")
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot parse user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := usersSrv.Update(id, user); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot update the user")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}
