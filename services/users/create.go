package users

import (
	"errors"
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
	uuid "github.com/satori/go.uuid"
)

// Create an user
func Create(user *models.User) error {
	db := pg.Connection()

	if !db.NewRecord(user) {
		return errors.New("User already exists")
	}

	tx := db.Begin()

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
