package services

import (
<<<<<<< HEAD
	"fmt"

=======
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/event"
)

//EventServicesImpl ... implements event.Event
<<<<<<< HEAD
type EventServicesImpl struct {
=======
type EventServicesImpl struct{
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	eventRepo event.EventRepository
}

//NewEventServiceImpl ... will create new EventService object
<<<<<<< HEAD
func NewEventServiceImpl(evntRepo event.EventRepository) *EventServicesImpl {
=======
func NewEventServiceImpl(evntRepo event.EventRepository) *EventServicesImpl{
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	return &EventServicesImpl{eventRepo: evntRepo}
}

//Events ... return list of events
<<<<<<< HEAD
func (esi *EventServicesImpl) Events() ([]entity.Event, []error) {
	evts, errs := esi.eventRepo.Events()
	if len(errs) > 0 {
		return nil, errs
	}
	return evts, nil
}

//Event ... reurn a specific event
func (esi *EventServicesImpl) Event(id uint) (*entity.Event, []error) {
	evnt, errs := esi.eventRepo.Event(id)
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println("--service--evnt returnd----", evnt)
	return evnt, errs

}

//UpdateEvent ... updates already exising event
func (esi *EventServicesImpl) UpdateEvent(event *entity.Event) (*entity.Event, []error) {
	evnt, errs := esi.eventRepo.UpdateEvent(event)
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs

}

//DeleteEvent ... delets an event
func (esi *EventServicesImpl) DeleteEvent(id uint) (*entity.Event, []error) {

	evnt, errs := esi.eventRepo.DeleteEvent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs

}

//AddEvent ... creates new event
func (esi *EventServicesImpl) AddEvent(event *entity.Event) (*entity.Event, []error) {

	evnt, errs := esi.eventRepo.AddEvent(event)
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs
}

//UpcomingEvents ... events that r not closed
func (esi *EventServicesImpl) UpcomingEvents() ([]entity.Event, []error) {
	evnt, errs := esi.eventRepo.UpcomingEvents()
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs
}

//GetUserSubscribedEvents ... returns user specific event(based on subscription/hobby)
func (esi *EventServicesImpl) GetUserSubscribedEvents(id uint) ([]entity.Event, error) {
	evnts, err := esi.eventRepo.GetUserSubscribedEvents(id)
	if err != nil {
		return nil, err
	}
	return evnts, nil
=======
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
//GetUserSubscribedEvents ... returns user specific event(based on subscription/hobby)
func(esi *EventServicesImpl) GetUserSubscribedEvents(id uint)([]entity.Event,error){
	evnts,err := esi.eventRepo.GetUserSubscribedEvents(id)
	if err != nil{
		return nil,err
	}
	return evnts,nil
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
}
