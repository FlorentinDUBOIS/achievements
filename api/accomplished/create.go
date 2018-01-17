package accomplished

import (
	"net/http"

	accomplishedSrv "github.com/FlorentinDUBOIS/achievements/services/accomplished"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// SetAccomplished an achievement
func SetAccomplished(ctx echo.Context) error {
	aid := ctx.Param("aid")
	uid := ctx.Param("uid")

	if err := accomplishedSrv.SetAccomplished(uid, aid); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"aid":   aid,
			"uid":   uid,
		}).Error("Cannot set the achievement accomplished")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
