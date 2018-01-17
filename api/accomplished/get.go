package accomplished

import (
	"net/http"

	accomplishedSrv "github.com/FlorentinDUBOIS/achievements/services/accomplished"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Get all achievements
func Get(ctx echo.Context) error {
	uid := ctx.Param("uid")

	accomplisheds, err := accomplishedSrv.Get(uid)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"uid":   uid,
		}).Error("Cannot retrieve all achievements")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(accomplisheds) == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, accomplisheds)
}
