package services

import (
	"github.com/goEventListing/event"
)

// Events() ([]entity.Event, []error) //get list of events
// Event(id uint) (*entity.Event, []error)
// //where isPassed is false
// UpcomingEvents() ([]entity.Event, []error)
// AddEvent(event *entity.Event)(*entity.Event, []error)
// UpdateEvent(event *entity.Event) (*entity.Event, []error)
// DeleteEvent(id uint) (*entity.Event,[]error)

//EventServicesImpl ... implements event.Event
type EventServicesImpl struct{
	eventRepo event.EventRepository
}

//NewEventServiceImpl ... will create new EventService object
func NewEventServiceImpl(evntRepo event.EventRepository) *EventServicesImpl{
	return &EventServicesImpl{eventRepo: evntRepo}
}

