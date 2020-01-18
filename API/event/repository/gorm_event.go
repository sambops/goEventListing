package repository

import (
	"fmt"

<<<<<<< HEAD
	"github.com/goEventListing/API/entity"
	"github.com/jinzhu/gorm"
=======
	"github.com/jinzhu/gorm"
	"github.com/goEventListing/API/entity"

	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *gorm.DB
}

// NewEventRepoImp will create an object of EventRepoImpl
<<<<<<< HEAD
func NewEventRepoImp(con *gorm.DB) *EventRepoImpl {
	return &EventRepoImpl{conn: con}
}

//Events ... returns all Events from the database
func (eri *EventRepoImpl) Events() ([]entity.Event, []error) {
	events := []entity.Event{}
	errs := eri.conn.Find(&events).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return events, errs
}

//Event ... returns events with specified id
func (eri *EventRepoImpl) Event(id uint) (*entity.Event, []error) {
	event := entity.Event{}
	errs := eri.conn.First(&event, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println("gorm---evnt returnd---", event)
	return &event, errs
}

//UpdateEvent updates already posted events with some modification...
func (eri *EventRepoImpl) UpdateEvent(event *entity.Event) (*entity.Event, []error) {
	evnt := event
	errs := eri.conn.Save(evnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs

}

//AddEvent ... adds new event to our event table
func (eri *EventRepoImpl) AddEvent(event *entity.Event) (*entity.Event, []error) {
	evet := event
	errs := eri.conn.Create(evet).GetErrors()
	if len(errs) > 0 {
		fmt.Println("check")
		return nil, errs
	}
	return evet, errs
}

//DeleteEvent ... delete an event recored from our event table with the given id
func (eri *EventRepoImpl) DeleteEvent(id uint) (*entity.Event, []error) {
	evnt, errs := eri.Event(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = eri.conn.Delete(evnt, evnt.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return evnt, errs
}

//UpcomingEvents ... returns events that are not yet closed
func (eri *EventRepoImpl) UpcomingEvents() ([]entity.Event, []error) {
	event := []entity.Event{}

	// Get all matched records
	errs := eri.conn.Where("is_passed = ?", "f").Find(&event).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return event, errs

}

//GetUserSubscribedEvents ... return events which are subscribed by the current user
func (eri *EventRepoImpl) GetUserSubscribedEvents(id uint) ([]entity.Event, error) {
=======
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
	errs :=eri.conn.Where("is_passed = ?","f").Find(&event).GetErrors()
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
	fmt.Println("check")
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
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	//userTag := entity.UserTag{}
	//eventTag := entity.EventTag{}
	actualEvent := []entity.Event{}
	//we select * from user_tag where user id =1 and we get the tag_id
	// tagidrows,err:= eri.conn.Raw("SELECT tag_id FROM user_tag WHERE userID = ?",id).Rows()
	// if err != nil{
	// 	return nil,err
	// }
	// defer tagidrows.Close()

	// for tagidrows.Next(){
	// 	eri.conn.ScanRows(tagidrows,&userTag)
	// 	//then select event_id from event_tag where tagID = the user_id we get from the above query
	// 	eventidrows,err := eri.conn.Raw("SELECT event_id FROM event_tag WHERE tagID = ?",userTag.TagID).Rows()
	// 	if err != nil{
	// 		return nil,err
	// 	}
	// 	defer eventidrows.Close()
	// 	for eventidrows.Next(){
	// 		eri.conn.ScanRows(eventidrows,&eventTag)
	// 		//then finally select from event where event_id = the result form the above query
<<<<<<< HEAD
	// 		event,err := eri.conn.Raw("SELECT * FROM events WHERE eventID = ?",eventTag.EventID).Rows()
=======
	// 		event,err := eri.conn.Raw("SELECT * FROM events WHERE eventID = ?",eventTag.EventID).Rows() 
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	// 		if err != nil{
	// 			return nil,err
	// 		}
	// 		defer event.Close()

	// 		for event.Next(){
	// 			eri.conn.ScanRows(event,&actualEvent)
	// 		}
	// 	}

	// }
<<<<<<< HEAD

	//err:= select * from events where eventID in(select eventID from event_tag where tage_id in(select tag_id from user_tag where user Id = ?, id) )
	event, err := eri.conn.Raw("SELECT * FROM events WHERE id in(SELECT event_id FROM event_tag WHERE tag_id IN (SELECT tag_id FROM user_tag WHERE user_id = ?))", id).Rows()
	if err != nil {
		fmt.Println("here i'm ")
		fmt.Println(err)
		return nil, err
	}

	for event.Next() {
		eri.conn.ScanRows(event, &actualEvent)
	}
	return actualEvent, err
}

//AddEventTags ... this store the event_id and tag_id to event_tag tabl
func (eri *EventRepoImpl) AddEventTags(etag *entity.EventTag) (*entity.EventTag, []error) {
	tagg := etag
	errs := eri.conn.Create(tagg).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tagg, errs

}
=======
	
//err:= select * from events where eventID in(select eventID from event_tag where tage_id in(select tag_id from user_tag where user Id = ?, id) )
event,err :=eri.conn.Raw("SELECT * FROM events WHERE id in(SELECT event_id FROM event_tag WHERE tag_id IN (SELECT tag_id FROM user_tag WHERE user_id = ?))",id).Rows()
if err != nil{
	fmt.Println("here i'm ")
	fmt.Println(err)
	return nil,err
}

for event.Next(){
	eri.conn.ScanRows(event,&actualEvent)
}
return actualEvent,err
}

//AddEventTags ... this store the event_id and tag_id to event_tag tabl
func(eri *EventRepoImpl) AddEventTags(etag *entity.EventTag)(*entity.EventTag,[]error){
	tagg := etag
errs := eri.conn.Create(tagg).GetErrors()
if len(errs)>0{
	return nil,errs
}
return tagg,errs
	 
}




>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
