package event

import (
	"github.com/EventListing/entity"
)

//USECASE
//this is our event usescase(has (interface)abstract classes that outer layers can use)
type EventService interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)
	UpcomingEvents() ([]entity.Event, error)
	
	Post(event entity.Event) error
	addTag(id []int)                           error//?? how do we add multiple tags
	UpdateEvent(event entity.Event) error
	
	DeleteEvent(id int) error
	rate(EventID, UserID, rating int)  error
	setRating(eventID int) error //this is done in the back
	
	getMyRating(UID, EventID int)    int	
	notify(eventID int, tagsID []int)          error
}
