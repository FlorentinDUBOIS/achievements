package accomplished

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
	rankSrv "github.com/FlorentinDUBOIS/achievements/services/rank"
)

// SetAccomplished an achievement
func SetAccomplished(uid, aid string) error {
	db := pg.Connection()
	now := time.Now()

	accomplished := &models.Accomplished{
		AchievementID: aid,
		UserID:        uid,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if db.Create(accomplished).Error != nil {
		return db.Error
	}

	count := 0
	db.
		Model(new(models.Accomplished)).
		Where("user_id = ?", uid).
		Count(&count)

	if db.Error != nil {
		return db.Error
	}

	rank := &models.Rank{
		UserID: uid,
		Count:  count,
	}

	if err := rankSrv.Save(rank); err != nil {
		return err
	}

	return rankSrv.Publish(rank)
}
