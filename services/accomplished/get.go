package accomplished

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Get all accomplished achievements
func Get(uid string) ([]models.Accomplished, error) {
	accomplisheds := []models.Accomplished{}
	db := store.Connection()

	db.
		Where("user_id = ?", uid).
		Select(accomplisheds)

	return accomplisheds, db.Error
}
