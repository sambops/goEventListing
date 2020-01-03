package handler

import (
	"html/template"
	"net/http"
	
	
	
)

type EventHandler struct {
	templ     *template.Template
	eventServ events.EventService
	
}

func NewEventHandler(T *template.Template, ES events.EventService) *EventHandler {
	return &EventHandler{templ: T, eventServ: ES}

}

