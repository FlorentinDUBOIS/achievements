package achievements

import (
	"net/http"

	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
)

// Delete an achievement
func Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := achievementsSrv.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
