package ranking

import (
	"net/http"

	rankSrv "github.com/FlorentinDUBOIS/achievements/services/rank"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Get the current ranking
func Get(ctx echo.Context) error {
	ranks, err := rankSrv.Get()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Cannot get ranks")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(ranks) == 0 {
		log.Info("No rank information found")
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, ranks)
}
