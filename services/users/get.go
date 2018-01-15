package users

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Get all users
func Get() ([]models.User, error) {
	db := store.Connection()
	users := []models.User{}

	db.Find(&users)
	if db.RecordNotFound() {
		return users, nil
	}

	return users, db.Error
}

// GetOne user
func GetOne(id string) (*models.User, error) {
	db := store.Connection()
	user := new(models.User)

	db.
		Where("id = ?", id).
		Limit(1).
		Find(user)

	if db.RecordNotFound() {
		return user, nil
	}

	return user, db.Error
}
