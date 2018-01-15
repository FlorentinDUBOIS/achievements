package achievements

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/models"
	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
)

// Update an achievement
func Update(ctx echo.Context) error {
	id := ctx.Param("id")
	achievement := new(models.Achievement)
	if err := ctx.Bind(achievement); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := achievementsSrv.Update(id, achievement); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, achievement)
}
