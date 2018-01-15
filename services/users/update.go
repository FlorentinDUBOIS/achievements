package users

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Update an user
func Update(id string, user *models.User) error {
	db := store.Connection()

	user.UpdatedAt = time.Now()

	db.
		Model(user).
		Where("users.id = ?", id).
		Updates(user)

	return db.Error
}
