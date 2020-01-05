package services

import (
	"github.com/goEventListing/entity"
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
//Events ... return list of events
func(esi *EventServicesImpl) Events() ([]entity.Event, []error){
	evts,errs := esi.eventRepo.Events()
	if len(errs) > 0{
		return nil,errs
	}
	return evts,nil
}

//Event ... reurn a specific event
func(esi *EventServicesImpl) Event(id uint) (*entity.Event, []error){
	evnt,errs := esi.eventRepo.Event(id)
	if len(errs) > 0{
		return nil,errs
	}
	return evnt,errs

}

//UpcomingEvents ... events that r not closed
func(esi *EventServicesImpl) UpcomingEvents() ([]entity.Event, []error){
evnt,errs := esi.eventRepo.UpcomingEvents()
if len(errs) > 0{
	return nil,errs
}
return evnt,errs
}

//AddEvent ... creates new event
func(esi *EventServicesImpl) AddEvent(event *entity.Event)(*entity.Event, []error){
evnt,errs := esi.eventRepo.AddEvent(event)
if len(errs) > 0{
	return nil,errs
}
return evnt,errs
}

//UpdateEvent ... updates already exising event
func(esi *EventServicesImpl) UpdateEvent(event *entity.Event) (*entity.Event, []error){
	evnt,errs:=esi.eventRepo.UpdateEvent(event)
	if len(errs) > 0 {
		return nil,errs
	}
	return evnt,errs

}
//DeleteEvent ... delets an event
func(esi *EventServicesImpl) DeleteEvent(id uint) (*entity.Event,[]error){
evnt,errs := esi.eventRepo.DeleteEvent(id)
if len(errs)>0 {
	return nil,errs
}
return evnt,errs

}
