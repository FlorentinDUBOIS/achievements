package accomplished

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// SetAccomplished an achievement
func SetAccomplished(uid, aid string) error {
	db := store.Connection()
	now := time.Now()

	db.
		Create(&models.Accomplished{
			AchievementID: aid,
			UserID:        uid,
			CreatedAt:     now,
			UpdatedAt:     now,
		})

	return db.Error
}
