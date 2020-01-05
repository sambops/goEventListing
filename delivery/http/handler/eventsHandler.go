package handler

import (
	"net/http"
	"github.com/goEventListing/event"
	"html/template"
)

//EventHandler handles event related requests
type EventHandler struct {
	templ     *template.Template
	eventServ event.EventServices
}

//NewEventHandler initializes and returns newEventHandler
func NewEventHandler(T *template.Template, ES event.EventServices) *EventHandler {
	return &EventHandler{templ: T, eventServ: ES}
}

//Index ... home page after login
func(eh *EventHandler) Index(w http.ResponseWriter,req *http.Request){
	
}
//AllEvents ... return all events form the datbase based on some condition(linke posted time)
func (eh *EventHandler) AllEvents(w http.ResponseWriter,req *http.Request){

}

//Event ... selects a specific event on select by event id
func (eh *EventHandler) Event(w http.ResponseWriter,req *http.Request){

}
//UpcomingEvents ... events that r not closed and event that r not colsed related to the user in the first order
func (eh *EventHandler) UpcomingEvents(w http.ResponseWriter,req *http.Request){

}

//CreateEvent ... creates new event
func (eh *EventHandler) CreateEvent(w http.ResponseWriter,req *http.Request){

}

//GetUserSpecificEvent ... display events based on the peron hobby
func (eh *EventHandler) GetUserSpecificEvent(w http.ResponseWriter,req *http.Request){

}





