package models

import "time"

// Accomplished structure
type Accomplished struct {
	UserID        string     `json:"user_id" gorm:"primary_key;type:uuid;not null"`
	AchievementID string     `json:"achievement_id" gorm:"primary_key;type:uuid;not null"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
