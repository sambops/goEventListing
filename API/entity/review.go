package entity

import "time"


// Review is when a user rates to an event
type Review struct {
	ID      uint 
	Rating  int `json:"rating"`
	EventID uint 
	UserID  uint  
	Event []Event `gorm:"many2many:user_event"`
	User []User `gorm:"many2many:user_review"`
	Message string `json:"message" gorm:"type:text;not null"`

	ReviewedAt time.Time
	// isempty    bool
}
