package database


import (
	"time"
)



//Event shows event intity
type Event struct {
	//gorm.Model //i use it to get when the event is CreatedAt
	ID uint `json:"id"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Details string `json:"details" gorm:"type:text;not null"`
	Country string `json:"country" gorm:"type:varchar(255)"`
	City string  `json:"city" gorm:"type:varchar(255)"`
	Place string `json:"place" gorm:"type:varchar(255)"`
	Price *float32 `json:"price" gorm:"type:numeric;not null;DEFAULT:0"`
	Image string `json:"image" gorm:"type:varchar(255)"`
	UserRefer uint //this is a forign key referencing USER
	TagRefer uint //this is a forign key referencing EVENTTAGE
	IsPassed   *bool `json:"ispassed" gorm:"type:bool;not null;DEFAULT:false"`
	Tag []Tag `gorm:"many2many:event_tag"`
	PlacedAt time.Time `json:"placedat"`
	
}


//Tag there are multiple tags for the user to choose from
type Tag struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"type:varchar(255);not null;unique"` 
	//description string `gorm:"type:text;not null"` 
	//icon        string `gorm:"type:varchar(255);not null"` 
	User []User `gorm:"many2many:user_tag"`
	Event []Event `gorm:"many2many:event_tag"`

}

//UserTag is struct containing event id and tag id
type UserTag struct{
	TagID uint
	UserID uint
}

//User ... represents users of our system
type User struct {
	ID uint 
	FirstName string `gorm:"type:varchar(64);not null"`
	LastName  string `gorm:"type:varchar(64);not null"`
	UserName  string `gorm:"type:varchar(64);not null;unique"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  []byte `gorm:"type:bytea;not null;unique"`
	Phone     string `gorm:"type:varchar(64);not null"`
	Image     string `gorm:"type:varchar(255);not null"`
	Event []Event `gorm:"foreignkey:UserRefer"`
	Tag []Tag `gorm:"many2many:user_tag"`
	PlacedAt time.Time
}

//EventTag is struct containing event id and tag id
type EventTag struct{
	TagID uint
	EventID uint
}