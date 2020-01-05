package handler

import (
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


