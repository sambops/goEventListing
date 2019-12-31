package entity

import "time"

//Event shows event intity
type Event struct {
	ID int

	Name, Details, Image              string
	UserID, CategoryID                int
	City, Country, Place, Coordinates string

	IsPassed   bool
	Rating     int
	PostedDate time.Time
	price      float32
}
