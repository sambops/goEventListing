package entity

import (
	"time"
)

//Event shows event intity
type Event struct {
	// gorm.Model

	ID       uint      `json:"id"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Details  string    `json:"details" gorm:"type:text;not null"`
	Image    string    `json:"image" gorm:"type:varchar(255)"`
	Price    *float32  `json:"price" gorm:"type:numeric;not null;DEFAULT:0"`
	PlacedAt time.Time `json:"placedat"`
	Rating   int

	Country string `json:"country" gorm:"type:varchar(255)"`
	City    string `json:"city" gorm:"type:varchar(255)"`
	Place   string `json:"place" gorm:"type:varchar(255)"`

	Coordinates string `json:"coordinates" gorm:"type:varchar(255)"`
	Date        string `json:"date" gorm:"type:varchar(255)"`

	UserRefer uint     //this is a forign key referencing USER
	TagRefer  uint     //this is a forign key referencing EVENTTAGE
	IsPassed  *bool    `json:"ispassed" gorm:"type:bool;not null;DEFAULT:false"`
	Tag       []Tag    `gorm:"many2many:event_tag"`
	Reviews   []Review `gorm:"many2many:event_Review"`
	user      []User   `gorm:"many2many:event_user"`
}

//User ... represents users of our system
type User struct {
	ID        uint
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	UserName  string  `json:"username"`
	Email     string  `json:"email"`
	Password  []byte  `json:"password"`
	Phone     string  `json:"phone"`
	Image     string  `json:"image"`
	Event     []Event `gorm:"foreignkey:UserRefer"`
	Tag       []Tag   `gorm:"many2many:user_tag"`
	PlacedAt  time.Time
	reviews   []Review `gorm:"many2many:user_review"`
}

// Review is when a user rates to an event
type Review struct {
	ID      uint   `json:"id"`
	Rating  int    `json:"rating"`
	EventID uint   `json:"event_id"`
	UserID  uint   `json:"user_id"`
	Message string `json:"details" gorm:"type:text;not null"`

	ReviewedAt time.Time
	// isempty    bool
}

// Tag there are multiple tags for the user to choose from
type Tag struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description string `json:"description" gorm:"type:text; not null"`
	Icon        string `json:"icon" gorm:"type:varchar(255);not null"`
}

// EventTag ...
type EventTag struct {
	TagID   uint
	EventID uint
}

// Notification whenbevent is posted it takes TagID from event_tag & match it with user_tags and inserts into the notification table
type Notification struct {
	ID      int
	EventID uint
	UserID  uint
	status  bool      //to show that it is seenor not there must be a tigger when the user opens it it will turn it to false
	EndDate time.Time //
}

//Authenticate ... used for Authentication() in user handler
type Authenticate struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// //UserTag is struct containing event id and tag id
// type UserTag struct{
// 	TagID uint
// 	UserID uint
// }
