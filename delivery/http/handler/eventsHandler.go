package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
	"github.com/birukbelay/Aprojects/goEventListing/event"
)

// EventHandler ...
type EventHandler struct {
	tmpl      *template.Template
	eventServ event.EventServices
}

// NewEventHandler ...
func NewEventHandler(T *template.Template, ES event.EventServices) *EventHandler {
	return &EventHandler{tmpl: T, eventServ: ES}

}

// Events ...
func (eh *EventHandler) Events(w http.ResponseWriter, r *http.Request) {
	events, errs := eh.eventServ.Events()
	if errs != nil {
		panic(errs)
	}
	eh.tmpl.ExecuteTemplate(w, "admin.categ.layout", events)
}

//Event ...
func (eh *EventHandler) Event(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		event := entity.Event{}
		event, errs := eh.eventServ.Event(id)

		if errs != nil {
			panic(errs)
		}

		eh.tmpl.ExecuteTemplate(w, "event.single", event)

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
