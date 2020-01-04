package event

import (
	"github.com/goEventListing/entity"
)

//EventServices  ... specify services related to event
type EventServices interface {
	Events() ([]entity.Event, []error) //get list of events
	Event(id uint) (*entity.Event, []error)

	//where isPassed is false
	UpcomingEvents() ([]entity.Event, []error)

	AddEvent(event *entity.Event)(*entity.Event, []error)
	//add event tags
	//AddEventTag(id []int)(*entity.Tag, []error) //?? how do we add multiple tags
	//notify(eventID uint, tagsID []int) []error //this should be done separatly in notification section
	//get the event tags
	//GetTags() ([]entity.Tag, []error)
	UpdateEvent(event *entity.Event) (*entity.Event, []error)
	DeleteEvent(id uint) (*entity.Event,[]error)

}
