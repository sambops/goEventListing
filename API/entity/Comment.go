package entity

import (
	"time"
)

//comment shows the comments forwarded by the users
type Comment struct {
	ID         uint      `json:"id"`
	Message    string    `json:"details" gorm:"type:text;not null"`
	UserRefer  uint      //this is a forign key referencing USER
	EventRefer uint      //this is a foriegn key referencing event
	PlacedAt   time.Time `json:"placedat"`
}
