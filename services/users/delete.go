package users

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Delete an user
func Delete(id string) error {
	db := store.Connection()

	db.
		Where("id = ?", id).
		Delete(new(models.User))

	return db.Error
}
