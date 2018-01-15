package achievements

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
)

// Get all achievements
func Get() ([]models.Achievement, error) {
	db := store.Connection()
	achievements := []models.Achievement{}

	db.Find(&achievements)
	if db.RecordNotFound() {
		return achievements, nil
	}

	return achievements, db.Error
}

// GetOne achievement
func GetOne(id string) (*models.Achievement, error) {
	db := store.Connection()
	achievement := new(models.Achievement)

	db.Where("id = ?", id).First(achievement)
	if db.RecordNotFound() {
		return achievement, nil
	}

	return achievement, db.Error
}