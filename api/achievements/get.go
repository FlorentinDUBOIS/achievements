package achievements

import (
	"net/http"

	achievementsSrv "github.com/FlorentinDUBOIS/achievements/services/achievements"
	"github.com/labstack/echo"
)

// Get all achievements
func Get(ctx echo.Context) error {
	achievements, err := achievementsSrv.Get()
	if err != nil {
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if achievement == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return ctx.JSON(http.StatusOK, achievement)
}
