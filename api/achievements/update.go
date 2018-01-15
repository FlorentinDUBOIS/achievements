package achievements

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/models"
	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Update an achievement
func Update(ctx echo.Context) error {
	id := ctx.Param("id")
	achievement := new(models.Achievement)
	if err := ctx.Bind(achievement); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot parse achievement")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := achievementsSrv.Update(id, achievement); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot update the achievement")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, achievement)
}
