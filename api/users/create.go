package users

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/models"
	usersSrv "github.com/FlorentinDUBOIS/achievements/services/users"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Create an user
func Create(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot parse user")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := usersSrv.Create(user); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot create user")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, user)
}
