package main

import (
	eventRepo"github.com/goEventListing/event/repository"
	eventService "github.com/goEventListing/event/services"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/goEventListing/delivery/http/handler"
	"github.com/goEventListing/user/services"
	_ "github.com/lib/pq"

	"html/template"

	"github.com/goEventListing/user/repository"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/goeventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()


	
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	//user
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(tmpl, userService)

	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(tmpl,eventService)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	

	http.HandleFunc("user/", userHandler.Index)
	http.HandleFunc("user/login", userHandler.Login)
	http.HandleFunc("user/register",userHandler.Register)
	http.HandleFunc("user/logout",userHandler.Logout)

	http.ListenAndServe(":8080", nil)

}
		