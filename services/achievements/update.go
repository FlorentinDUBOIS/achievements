package achievements

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Update an achievement
func Update(id string, achievement *models.Achievement) error {
	db := pg.Connection()

	achievement.UpdatedAt = time.Now()

	db.
		Model(achievement).
		Where("achievements.id = ?", id).
		Updates(achievement)

	return db.Error
}
