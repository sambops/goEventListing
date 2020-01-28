package handler

import (
	"io"
	"path/filepath"
	"os"
	"mime/multipart"
	"github.com/goEventListing/client/form"
	"net/url"
	"github.com/goEventListing/client/rtoken"
	//"fmt"
	"strconv"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"net/http"
	"html/template"
)

//EventHandler ... handles event related requests
type EventHandler struct{
	tmpl   *template.Template
	csrfSignKey []byte
	ushnd   UserHandler
}

//NewEventHandler ... initializes and returns new EventHandler
func NewEventHandler(T *template.Template,csKey []byte) *EventHandler{
	return &EventHandler{tmpl: T,csrfSignKey: csKey}
}



//Events handle reques on route/events
func(eh *EventHandler) Events(w http.ResponseWriter,req *http.Request){
	
	evt,err := service.AllEvents()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	token, err := rtoken.CSRFToken(eh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Events 		*[]entity.Event
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Events: 	evt,
		CSRF:       token,
	}

	// tmplData := struct {
	// 	event  *[]entity.Event
	// 	user   *entity.User
		
	// }{
	// 	event:  evt,
	// 	user: usr,
	// }
	eh.tmpl.ExecuteTemplate(w, "all.layout", tmplData)

}

//Upcoming handle request on route/upcoming
func(eh *EventHandler) Upcoming(w http.ResponseWriter,req *http.Request){
	// //usr := GetUser(w,req)
	// usr, errr := service.GetUser(usr.ID)

	// if errr != nil {
	// 		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	// 	}
	upcoming,err := service.UpcomingEvent()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(eh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	// tmplData := struct {
	// 	event  *[]entity.Event
	// 	user   *entity.User
		
	// }{
	// 	event:  evt,
	// 	user: usr,
	// }
	
	// if err != nil{
	// 	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	// }
	// eh.tmpl.ExecuteTemplate(w,"upcomingevent.html",tmplData)
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Upcoming *[]entity.Event
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Upcoming: 	upcoming,
		CSRF:       token,
	}
	eh.tmpl.ExecuteTemplate(w, "upcoming.html", tmplData)
}
//CreateEvent ... request on route/create
func(eh *EventHandler) CreateEvent(w http.ResponseWriter,req *http.Request){
 	id := eh.ushnd.loggedInUser.ID
	
	token, err := rtoken.CSRFToken(eh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if req.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		eh.tmpl.ExecuteTemplate(w, "/", newCatForm)
	}
	if req.Method == http.MethodPost {
		// Parse the form data
		err := req.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		newEvtForm := form.Input{Values: req.PostForm, VErrors: form.ValidationErrors{}}
		newEvtForm.Required("name", "details","country","city","place","price")
		newEvtForm.MinLength("details", 15)
		newEvtForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newEvtForm.Valid() {
			eh.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newEvtForm)
			return
		}
		
		// mf, fh, err := req.FormFile("image")
		// if err != nil {
		// 	newEvtForm.VErrors.Add("image", "File error")
		// 	eh.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newEvtForm)
		// 	return
		// }
		//defer mf.Close()

		price,_ := strconv.ParseFloat(req.FormValue("price"),64)

		evt := &entity.Event{
			Name:        req.FormValue("name"),
			Details : req.FormValue("details"),
			Country : req.FormValue("country"),
			City : req.FormValue("city"),
			Place : req.FormValue("place"),
			Price : &price,
			UserID :id,
			Image:  "img.jpg",
			//fh.Filename,
		}
		//writeFile(&mf, fh.Filename)
		//_, errs := ach.categorySrv.StoreCategory(ctg)
		_,err = service.AddEvent(evt)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

//UserSpecific handle request on route/upcoming
func (eh *EventHandler) UserSpecific(w http.ResponseWriter,req *http.Request){
	id := eh.ushnd.loggedInUser.ID

	
	// eh.tmpl.ExecuteTemplate(w,"foru.html",tmplData)
	evnts,err := service.GetUserSubscribedEvents(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(eh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		UserSpecific *[]entity.Event
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		UserSpecific: evnts,
		CSRF:       token,
	}
	eh.tmpl.ExecuteTemplate(w, "foru.html", tmplData)
}

//RemoveEvent ... handle request on route/remove
func (eh *EventHandler) RemoveEvent(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_,err = service.DeleteEvent(uint (id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}




func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
