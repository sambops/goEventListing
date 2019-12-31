package events

import "github.com/birukbelay/Aprojects/EventListing/entity"

// EventServices repository
type EventServices interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)
	UpcomingEvents() ([]entity.Event, error)

	Post(event entity.Event) error
	addTag(id []int) error //?? how do we add multiple tags
	UpdateEvent(event entity.Event) error

	DeleteEvent(id int) error

	notify(eventID int, tagsID []int) error
}
