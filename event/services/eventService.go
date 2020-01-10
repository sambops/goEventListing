package services

import (
	en "github.com/birukbelay/Aprojects/goEventListing/entity"
	eve "github.com/birukbelay/Aprojects/goEventListing/event"
	// "github.com/birukbelay/Aprojects/goEventListing/entity"
)

// EventServiceImpl implements review.Reviewservice interface
type EventServicesImpl struct {
	EventRepo eve.EventRepository
}

//NewEventServicesImpl ... will create new EventService object
func NewEventServicesImpl(evntRepo eve.EventRepository) *EventServicesImpl {
	return &EventServicesImpl{EventRepo: evntRepo}
}

func (es *EventServicesImpl) Events() ([]en.Event, error) {

	events, err := es.EventRepo.Events()

	if err != nil {
		return nil, err
	}

	return events, nil
}

//Event ... return list of events
func (esi *EventServicesImpl) Event(id int) (en.Event, error) {

	evnt, err := esi.EventRepo.Event(id)
	if err != nil {
		return evnt, err
	}
	return evnt, nil
}
