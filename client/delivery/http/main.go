package main

import (
	"github.com/goEventListing/client/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
)


var tmpl = template.Must(template.ParseGlob("../../ui/templates/*.html"))

func main() {
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router := httprouter.New()
	// http.HandleFunc("/", userHandler.Index)
	// http.HandleFunc("user/login", userHandler.Login)
	// http.HandleFunc("user/register",userHandler.Register)
	// http.HandleFunc("user/logout",userHandler.Logout)

	//user
	usrHandler := handler.NewUserHandler(tmpl)
	router.GET("/",usrHandler.Index)
	router.POST("/el/user/register",usrHandler.Register)
	router.POST("/el/user/login",usrHandler.Login)
	router.GET("/el/user/logout",usrHandler.Logout)

	//event
	evtHandler := handler.NewEventHandler(tmpl)
	router.GET("/el/event/all",evtHandler.Events)
	router.GET("/el/event/upcoming",evtHandler.Upcoming)
	router.GET("/el/event/create",evtHandler.CreateEvent)
	router.GET("/el/event/foru",evtHandler.UserSpecific)


	http.ListenAndServe(":8080",nil)
}