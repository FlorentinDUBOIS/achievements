package models

// Rank is the number of accomplished achievement of a user
type Rank struct {
	UserID string `json:"user_id"`
	Count  int    `json:"count"`
}
