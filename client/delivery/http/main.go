package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goEventListing/client/delivery/http/handler"
	"github.com/julienschmidt/httprouter"
)

var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {
	fs := http.FileServer(http.Dir("../../ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router := httprouter.New()
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
