package users

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Update an user
func Update(id string, user *models.User) error {
	db := pg.Connection()

	user.UpdatedAt = time.Now()

	db.
		Model(user).
		Where("users.id = ?", id).
		Updates(user)

	return db.Error
}
