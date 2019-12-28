package entity

//Event ... represents event
type Event struct {
	EventID  int
	Name     string
	Detail   string
	Location string // just for the time sake we take location in string
	Image    string
	State    bool
}
