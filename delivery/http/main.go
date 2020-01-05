package main

import (
	eventRepo"github.com/goEventListing/API/event/repository"
	eventService "github.com/goEventListing/API/event/services"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/goEventListing/API/delivery/http/handler"
	"github.com/goEventListing/API/user/services"
	_ "github.com/lib/pq"

	"html/template"

	"github.com/goEventListing/API/user/repository"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/goeventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()


	
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
		
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//user
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(tmpl, userService)


	
	http.HandleFunc("/", userHandler.Index)
	http.HandleFunc("user/login", userHandler.Login)
	http.HandleFunc("user/register",userHandler.Register)
	http.HandleFunc("user/logout",userHandler.Logout)


	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(tmpl,eventService)

	http.HandleFunc("event/allevents",eventHandler.AllEvents)
	http.HandleFunc("event/event",eventHandler.Event)
	http.HandleFunc("event/upcoming",eventHandler.UpcomingEvents)
	http.HandleFunc("event/create",eventHandler.CreateEvent)
	http.HandleFunc("event/foru",eventHandler.GetUserSpecificEvent)

	
	

	





	http.ListenAndServe(":8080", nil)

}
		