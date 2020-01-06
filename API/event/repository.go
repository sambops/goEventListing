package event

import (
	"github.com/goEventListing/API/entity"
)

// EventRepository repository
type EventRepository interface {
	Events() ([]entity.Event, []error) //get list of events
	Event(id uint) (*entity.Event, []error)
	//where isPassed is false
	UpcomingEvents() ([]entity.Event, []error)
	AddEvent(event *entity.Event)(*entity.Event, []error)
	UpdateEvent(event *entity.Event) (*entity.Event, []error)
	DeleteEvent(id uint) (*entity.Event,[]error)
	//user specific events
	GetUserSubscribedEvents(id uint)([]entity.Event,error)
}
