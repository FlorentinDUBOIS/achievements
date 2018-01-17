package accomplished

import (
	"net/http"

	accomplishedSrv "github.com/FlorentinDUBOIS/achievements/services/accomplished"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Delete an accomplished achievement
func Delete(ctx echo.Context) error {
	aid := ctx.Param("aid")
	uid := ctx.Param("uid")

	if err := accomplishedSrv.Delete(uid, aid); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"aid":   aid,
			"uid":   uid,
		}).Error("Cannot delete the accomplished achievement")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
