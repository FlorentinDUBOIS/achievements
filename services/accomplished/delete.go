package accomplished

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Delete an accomplished achievement
func Delete(uid, aid string) error {
	db := store.Connection()

	db.
		Where("user_id = ? and achievement_id = ?", uid, aid).
		Delete(&models.Accomplished{})

	return db.Error
}
