package event

import "github.com/birukbelay/Aprojects/goEventListing/entity"

// EventRepository repository
type EventRepository interface {
	Events() ([]entity.Event, error) //get list of events
	Event(id int) (entity.Event, error)
}
