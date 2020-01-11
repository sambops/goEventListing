package handler

import (
	"fmt"
	"strconv"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
)

//EventHandler ... handlers event related requests
type EventHandler struct{
	tmpl *template.Template
}

//NewEventHandler ... initializes and returns new EventHandler
func NewEventHandler(T *template.Template) *EventHandler{
	return &EventHandler{tmpl:T}
}



//Events handle reques on route/events
func(eh *EventHandler) Events(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	fmt.Println("kiki")
	usr := GetUser(w,req)
	usr, errr := service.GetUser(usr.ID)
	if errr != nil {
			// http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	
		}
	evt,err := service.AllEvents()

	tmplData := struct {
		event  *[]entity.Event
		user   *entity.User
		
	}{
		event:  evt,
		user: usr,
	}


	if err != nil{
		println("check")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	eh.tmpl.ExecuteTemplate(w,"allevent.html",tmplData)
}

//Upcoming handle request on route/upcoming
func(eh *EventHandler) Upcoming(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	usr := GetUser(w,req)
	usr, errr := service.GetUser(usr.ID)

	if errr != nil {
			http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		}

	evt,err := service.UpcomingEvent()

	tmplData := struct {
		event  *[]entity.Event
		user   *entity.User
		
	}{
		event:  evt,
		user: usr,
	}
	
	if err != nil{
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	eh.tmpl.ExecuteTemplate(w,"upcomingevent.html",tmplData)
}
//CreateEvent ... request on route/create
func(eh *EventHandler) CreateEvent(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	
	var evt *entity.Event
	if req.Method == http.MethodPost{
		// name := req.FormValue("name")
		// detail := req.FormValue("details")
		// country := req.FormValue("country")
		// city := req.FormValue("city")
		// place := req.FormValue("place")
		// price := req.FormValue("price")
		// img := req.FormValue("img")	
		
		// price,_ = strconv.ParseFloat(price,32)
		// evt = &entity.Event{Name:name,Details:detail,Country:country,City:city,Place:place,Price:price,Image:imag}

		// evt,err := service.AddEvent(evt)
		//redirect 
		http.Redirect(w,req,"/",http.StatusSeeOther)
		return
	}
	eh.tmpl.ExecuteTemplate(w,"createEvent.html",evt)
}

//UserSpecific handle request on route/upcoming
func (eh *EventHandler) UserSpecific(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	idraw := req.FormValue("id")
	id, err := strconv.Atoi(idraw)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}

	
	usr := GetUser(w,req)
	usr, errr := service.GetUser(usr.ID)

	if errr != nil {
			http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		}
		
	evt,err := service.GetUserSubscribedEvents(uint (id))	
	if err != nil{
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}

	tmplData := struct {
		event  *[]entity.Event
		user   *entity.User
		
	}{
		event:  evt,
		user: usr,
	}

	eh.tmpl.ExecuteTemplate(w,"foru.html",tmplData)
}

