package entity

import "time"

// Review is when a user rates to an event
type Review struct {
	ID              int
	Rating     int
	EventID int
	UserID int
	Message            string

	
	ReviewedAt time.Time
	// isempty    bool
}
