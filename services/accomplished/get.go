package accomplished

import (
	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Get all accomplished achievements
func Get(uid string) ([]models.Accomplished, error) {
	accomplisheds := []models.Accomplished{}
	db := pg.Connection()

	db.
		Where("user_id = ?", uid).
		Select(accomplisheds)

	return accomplisheds, db.Error
}
