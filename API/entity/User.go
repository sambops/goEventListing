package entity

import (
	"time"
	
)

//User ... represents users of our system
type User struct {
	ID uint
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	Phone     string `json:"phone"`
	Image     string `json:"image"`
	Event []Event `gorm:"foreignkey:UserRefer"`
	Tag []Tag `gorm:"many2many:user_tag"`
	Review []Review `gorm:"many2many:user_review"`
	PlacedAt time.Time
}
// CREATE TABLE users(
// 	user_id serial PRIMARY KEY,
// 	username VARCHAR (50) UNIQUE NOT NULL,
// 	first_name VARCHAR (50) NOT NULL,
// 	last_name VARCHAR (50) NOT NULL,
// 	email VARCHAR (50) UNIQUE NOT NULL,
// 	password TEXT NOT NULL,
// 	phone VARCHAR (50),
// 	image TEXT
//  );