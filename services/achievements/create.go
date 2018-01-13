package achievements

import (
	"errors"
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
	uuid "github.com/satori/go.uuid"
)

// Create an achievement
func Create(achievement *models.Achievement) error {
	db := store.Connection()

	if !db.NewRecord(achievement) {
		return errors.New("Achievement already exists")
	}

	tx := db.Begin()

	achievement.ID = uuid.NewV4().String()
	achievement.CreatedAt = time.Now()
	achievement.UpdatedAt = time.Now()

	if err := tx.Create(achievement).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
