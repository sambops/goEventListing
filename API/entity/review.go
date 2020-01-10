package entity

import "time"

// Review is when a user rates to an event
type Review struct {
	ID      uint
	Rating  int
	EventID int
	UserID  int
	Message string `json:"details" gorm:"type:text;not null"`

	ReviewedAt time.Time
	// isempty    bool
}
