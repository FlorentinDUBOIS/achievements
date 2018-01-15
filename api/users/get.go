package users

import (
	"net/http"

	usersSrv "github.com/FlorentinDUBOIS/achievements/services/users"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Get all users
func Get(ctx echo.Context) error {
	users, err := usersSrv.Get()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get all users")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(users) == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, users)
}

// GetOne user
func GetOne(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := usersSrv.GetOne(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get one user")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, user)
}
