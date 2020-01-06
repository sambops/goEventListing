package event

import (
	"github.com/goEventListing/entity"
)

//EXTERNALINTERFACE(DATABASE)
//EventRepository repository(interface)
type EventRepository interface {
	Events() ([]entity.Event, error)
	AddEvent(event entity.Event) error
	EditEvent(event entity.Event) error
	DeleteEvent(id int) error
}
