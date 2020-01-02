<<<<<<< HEAD
package events

import "github.com/birukbelay/Aprojects/EventListing/entity"

// EventRepository repository
type EventRepository interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)
	UpcomingEvents() ([]entity.Event, error)

	Post(event entity.Event) error
	addTag(id []int) error //?? how do we add multiple tags
	UpdateEvent(event entity.Event) error

	DeleteEvent(id int) error

	notify(eventID int, tagsID []int) error
=======
package event

import (
	"github.com/EventListing/entity"
)

//EXTERNALINTERFACE(DATABASE)
//EventRepository repository(interface)
type EventRepository interface {
	Events() ([]entity.Event, error)
	AddEvent(event entity.Event) error
	EditEvent(event entity.Event) error
	DeleteEvent(id int) error
>>>>>>> e18614362e5300c66820568db16e16c72c4c3f76
}
