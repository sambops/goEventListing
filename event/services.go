package event

import (
	"github.com/birukbelay/Aprojects/goEventListing/entity"
)

//USECASE
type EventServices interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)
}
