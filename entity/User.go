package entity

import (
	"github.com/jinzhu/gorm"
)

//User ... represents users of our system
type User struct {
	gorm.Model
	ID    int 
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  []byte
	Phone     string
	Image     string
	Event []Event `gorm:"foreignkey:UserRefer"`
	Tag []Tag `gorm:"many2many:user_tag"`
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