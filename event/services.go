package event

import (
	"github.com/goEventListing/entity"
)

//USECASE
//this is our event usescase(has (interface)abstract classes that outer layers can use)
type EventService interface {
	Events() ([]entity.Event, error)
	AddEvent(event entity.Event) error
	EditEvent(event entity.Event) error
	DeleteEvent(id int) error
}
