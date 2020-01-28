package entity

import "time"

// User represents user data
var MockUser = User{
	ID:        1,
	FirstName: "Emnet",
	LastName:  "Alazar",
	UserName:  "someone",
	Email:     "emnet@gmail.com",
	Password:  "secret",
	Phone:     "0926100732",
	Image:     "picture.PNG",
	RoleID:    1,
	Event:     nil,
	Review:    nil,
	Tag:       nil,
	PlacedAt: time.Date(2019, 12,12,10,0, 0,0, time.UTC),
}
