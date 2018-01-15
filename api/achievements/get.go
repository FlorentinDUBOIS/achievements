package achievements

import (
	"net/http"

	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Get all achievements
func Get(ctx echo.Context) error {
	achievements, err := achievementsSrv.Get()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get all achievements")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(achievements) == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, achievements)
}

// GetOne achievement
func GetOne(ctx echo.Context) error {
	id := ctx.Param("id")
	achievement, err := achievementsSrv.GetOne(id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get achievement")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if achievement == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, achievement)
}
