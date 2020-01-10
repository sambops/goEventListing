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

	Coordinates string `json:"city" gorm:"type:varchar(255)"`
	Date        string `json:"place" gorm:"type:varchar(255)"`

	UserRefer uint     //this is a forign key referencing USER
	TagRefer  uint     //this is a forign key referencing EVENTTAGE
	IsPassed  *bool    `json:"ispassed" gorm:"type:bool;not null;DEFAULT:false"`
	Tag       []Tag    `gorm:"many2many:event_tag"`
	Reviews   []Review `gorm:"many2many:event_Review"`
	user      []User   `gorm:"many2many:event_user"`
}
