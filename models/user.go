package models

import "time"

// User describe an startup week-end user
type User struct {
	ID        string     `json:"id"    gorm:"primary_key;type:uuid"`
	SWID      string     `json:"sw_id" gorm:"unique;not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
