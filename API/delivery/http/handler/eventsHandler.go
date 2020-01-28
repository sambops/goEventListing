package handler

import (
	"fmt"
	
	"github.com/goEventListing/API/entity"
	"strconv"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/goEventListing/API/event"
	"net/http"
	
)

//EventHandler handles event related requests
type EventHandler struct {
	eventServ event.EventServices
}

//NewEventHandler initializes and returns newEventHandler
func NewEventHandler(ES event.EventServices) *EventHandler {
	return &EventHandler{eventServ: ES}
}

//Index ... home page after login
func(eh *EventHandler) Index(w http.ResponseWriter,req *http.Request, _ httprouter.Params){
	
}

//AllEvents ... handles GET /event/allevents request
func (eh *EventHandler) AllEvents(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
evnt,err := eh.eventServ.Events()

if err != nil{
	w.Header().Set("Content-Type","application/json")
	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	return
}
output,errr := json.MarshalIndent(evnt,"","\t\t")
if errr!=nil{
	w.Header().Set("Content-Type","application/json")
	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
	return
}
w.Header().Set("Content-Type","application/json")
w.Write(output)
return

}

//Event ... handles GET /event/event/:id request
func (eh *EventHandler) Event(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))
	if err != nil{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	evnt,errs := eh.eventServ.Event(uint(id))
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(evnt,"","\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
	

}
//UpcomingEvents ...  handles GET /event/upcoming request
func (eh *EventHandler) UpcomingEvents(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	evnts,errs := eh.eventServ.UpcomingEvents()
	if len(errs) > 0{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(evnts,"","\t\t")
	if err!=nil{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
	
}

//CreateEvent ...  handles GET /event/create request
func (eh *EventHandler) CreateEvent(w http.ResponseWriter,req *http.Request,_ httprouter.Params){
	l :=req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)
	event := &entity.Event{}

	err:= json.Unmarshal(body,event)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(event)
	event,errs := eh.eventServ.AddEvent(event)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	p := fmt.Sprintf("/event/create/%d", event.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)

	output,err:= json.MarshalIndent(event,"","\t\t")
	if err != nil{
		
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(output)

	return
}

//GetUserSpecificEvent ... handles GET /event/foru request
func (eh *EventHandler) GetUserSpecificEvent(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))

	if err != nil{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	events,err := eh.eventServ.GetUserSubscribedEvents(uint (id))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(events,"","\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return

}
//UpdateEvent ...  handles GET /event/create request
func (eh *EventHandler) UpdateEvent(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))
	
if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
event,errs := eh.eventServ.Event(uint(id))
if len(errs) > 0{
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
l:= req.ContentLength
body := make([]byte,l)
req.Body.Read(body)

json.Unmarshal(body, &event)

event,errs = eh.eventServ.UpdateEvent(event)
if len(errs) > 0 {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
output, err := json.MarshalIndent(event, "", "\t\t")

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
w.Header().Set("Content-Type", "application/json")
w.Write(output)
return
}

//RemoveEvent ... handle POST /el/event/remove:id
func (eh *EventHandler) RemoveEvent(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_,errs := eh.eventServ.DeleteEvent(uint(id))
	
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

