package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/goEventListing/entity"

	
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *gorm.DB
	
}

// NewEventRepoImp will create an object of EventRepoImpl
func NewEventRepoImp(con *gorm.DB) *EventRepoImpl{
	return &EventRepoImpl{conn : con}
}

//Events ... returns all Events from the database
func (eri *EventRepoImpl) Events() ([]entity.Event, []error){
	events := []entity.Event{}
	errs := eri.conn.Find(&events).GetErrors()
	if len(errs) > 0 {
		return nil,errs
	}
	return events,errs
}
//Event ... returns events with specified id
func (eri *EventRepoImpl) Event(id uint)(*entity.Event,[]error){
	event :=entity.Event{}
	errs := eri.conn.First(&event,id).GetErrors()
	if len(errs)>0 {
		return nil,errs
	}
	return &event,errs
}
//UpcomingEvents ... returns events that are not yet closed
func (eri *EventRepoImpl) UpcomingEvents() ([]entity.Event, []error){
	event := []entity.Event{}
// Get all matched records
	errs :=eri.conn.Where("ispassed = ?","false").Find(&event).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return event,errs
	
}
//AddEvent ... adds new event to our event table
func (eri *EventRepoImpl) AddEvent(event *entity.Event)(*entity.Event, []error){
evet := event
errs := eri.conn.Create(evet).GetErrors()
if len(errs)>0{
	return nil,errs
}
return evet,errs
}

//UpdateEvent updates already posted events with some modification...
func (eri *EventRepoImpl) UpdateEvent(event *entity.Event) (*entity.Event, []error){
	evnt := event
	errs :=eri.conn.Save(evnt).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return evnt,errs

}
//DeleteEvent ... delete an event recored from our event table with the given id
func (eri *EventRepoImpl) DeleteEvent(id uint) (*entity.Event,[]error){
	evnt ,errs := eri.Event(id)
	if len(errs) > 0 {
		return nil,errs
	}
	errs = eri.conn.Delete(evnt,evnt.ID).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return evnt,errs
}



