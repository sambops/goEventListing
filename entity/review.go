package entity

import "time"

// Review is when a user rates to an event
type Review struct {
	ID              int
	UserID, EventID int
	Body            string
	rating          int
	reviewedAt         time.Time
}



