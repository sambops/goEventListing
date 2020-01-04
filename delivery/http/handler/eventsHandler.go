package handler

import (
	"html/template"

	"github.com/goEventListing/event"
)

type EventHandler struct {
	templ     *template.Template
	eventServ events.EventService
}

func NewEventHandler(T *template.Template, ES events.EventService) *EventHandler {
	return &EventHandler{templ: T, eventServ: ES}

}
