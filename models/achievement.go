package models

import "time"

// Achievement describe an achievement
type Achievement struct {
	ID          string     `json:"id"          gorm:"primary_key;type:uuid"`
	Name        string     `json:"name"        gorm:"unique;not null;type:varchar(256)"`
	Description string     `json:"description" gorm:"not null;type:varchar(2048)"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
