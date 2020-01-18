package main

import (
<<<<<<< HEAD
	"fmt"
	"html/template"
	"net/http"

	"github.com/goEventListing/client/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
)

=======
	"github.com/goEventListing/client/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
)


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
var tmpl = template.Must(template.ParseGlob("../../ui/templates/*.html"))

func main() {
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router := httprouter.New()
<<<<<<< HEAD
	eventhandler := handler.NewEventHandler(tmpl)
	router.GET("/event/:id", eventhandler.Event)

	usrHandler := handler.NewUserHandler(tmpl)
	router.GET("/", usrHandler.Index)
	router.POST("/el/user/register", usrHandler.Register)
	router.POST("/el/user/login", usrHandler.Login)
	router.GET("/el/user/logout", usrHandler.Logout)

	fmt.Println("...8082...")
	err := http.ListenAndServe(":8082", router)
	panic(err)
}
=======
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
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
