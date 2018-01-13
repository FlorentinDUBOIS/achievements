package achievements

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/models"
	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Create an achievement
func Create(ctx echo.Context) error {
	achievement := new(models.Achievement)
	if err := ctx.Bind(achievement); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot parse achievement")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := achievementsSrv.Create(achievement); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot create achievement")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, achievement)
}
