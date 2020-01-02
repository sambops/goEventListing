package entity

<<<<<<< HEAD
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
=======
//Event ... represents event
type Event struct {
	EventID  int
	Name     string
	Detail   string
	Location string // just for the time sake we take location in string
	Image    string
	State    bool
>>>>>>> e18614362e5300c66820568db16e16c72c4c3f76
}
