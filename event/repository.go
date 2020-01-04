package event

import (
	"github.com/goEventListing/entity"
)

// EventRepository repository
type EventRepository interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)

	UpcomingEvents() ([]entity.Event, error)

	getTags() ([]entity.Tag, error)

	Post(event entity.Event) error
	addTag(id []int) error //?? how do we add multiple tags

	notify(eventID int, tagsID []int) error

	UpdateEvent(event entity.Event) error
	DeleteEvent(id int) error
}
