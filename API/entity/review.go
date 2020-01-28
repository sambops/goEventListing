package entity

import "time"


// Review is when a user rates to an event
type Review struct {
	ID      uint `json:"id"`
	Rating  int `json:"rating"`
	
	User User 
	UserID	uint `json:"userid"` // forign key referencing User

	Event Event 
	EventID uint `json:"eventid"`// forign key referencing Event
	

	Message string `json:"message" gorm:"type:text;not null"`
	ReviewedAt time.Time
	// isempty    bool
}
