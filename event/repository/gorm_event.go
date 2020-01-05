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
//GetUserSubscribedEvents ... return events which are subscribed by the current user
func(eri *EventRepoImpl) GetUserSubscribedEvents(id uint)([]entity.Event,error){
	userTag := entity.UserTag{}
	eventTag := entity.EventTag{}
	actualEvent := []entity.Event{}
	//we select * from user_tag where user id =1 and we get the tag_id
	tagidrows,err:= eri.conn.Raw("SELECT tag_id FROM user_tag WHERE userID = ?",id).Rows()
	if err != nil{
		return nil,err
	}
	defer tagidrows.Close()

	for tagidrows.Next(){
		eri.conn.ScanRows(tagidrows,&userTag)
		//then select event_id from event_tag where tagID = the user_id we get from the above query
		eventidrows,err := eri.conn.Raw("SELECT event_id FROM event_tag WHERE tagID = ?",userTag.TagID).Rows()
		if err != nil{
			return nil,err
		}
		defer eventidrows.Close()
		for eventidrows.Next(){
			eri.conn.ScanRows(eventidrows,&eventTag)
			//then finally select from event where event_id = the result form the above query
			event,err := eri.conn.Raw("SELECT * FROM events WHERE eventID = ?",eventTag.EventID).Rows() 
			if err != nil{
				return nil,err
			}
			defer event.Close()

			for event.Next(){
				eri.conn.ScanRows(event,&actualEvent)
			}
		}

	}
	return actualEvent,err

}
//



