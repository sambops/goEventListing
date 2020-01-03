package event

import (
<<<<<<< HEAD
	"github.com/goEventListing/entity"
=======
	"github.com/birukbelay/Aprojects/goEventListing/entity"
>>>>>>> 88892e25e6fb176ea4f4ccd0ea65914deaca56f9
)

//USECASE
type EventServices interface {
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
