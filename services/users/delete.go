package users

import (
	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Delete an user
func Delete(id string) error {
	db := pg.Connection()

	db.
		Where("id = ?", id).
		Delete(new(models.User))

	return db.Error
}
