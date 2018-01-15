package achievements

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Delete an achievement
func Delete(id string) error {
	db := store.Connection()

	db.
		Where("id = ?", id).
		Delete(models.Achievement{})

	return db.Error
}
