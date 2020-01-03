package handler

import (
	"github.com/goEventListing/event"
	"html/template"

	
	
	
)

type EventHandler struct {
	templ     *template.Template
	eventServ event.EventServices
	
	
}

func NewEventHandler(T *template.Template, ES event.EventServices) *EventHandler {
	return &EventHandler{templ: T, eventServ: ES}

}

