package achievements

import (
	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Delete an achievement
func Delete(id string) error {
	db := pg.Connection()

	db.
		Where("id = ?", id).
		Delete(models.Achievement{})

	return db.Error
}
