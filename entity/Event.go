package entity

import (
	"github.com/jinzhu/gorm"

)

//Event shows event intity
type Event struct {
	gorm.Model //i use it to get when the event is CreatedAt
	ID uint  
	Name string `gorm:"type:varchar(255);not null"`
	Details string `gorm:"type:text;not null"`
	Country string `gorm:"type:varchar(255)"`
	City string `gorm:"type:varchar(255)"`
	Place string `gorm:"type:varchar(255)"`
	Price *float32 `gorm:"type:numeric;not null;DEFAULT:0"`
	Image string `gorm:"type:varchar(255)"`
	UserRefer uint //this is a forign key referencing USER
	TagRefer uint //this is a forign key referencing EVENTTAGE
	IsPassed   *bool `gorm:"type:bool;not null;DEFAULT:false"`
	Tag []Tag `gorm:"many2many:event_tag"`
	
}


