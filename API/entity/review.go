package entity

import "time"


// Review is when a user rates to an event
type Review struct {
	ID      uint 
	Rating  int `json:"rating"`
	UserRefer uint // forign key referencing User
	EventRefer uint // forign key referencing Event
	Message string `json:"message" gorm:"type:text;not null"`
	ReviewedAt time.Time
	// isempty    bool
}
