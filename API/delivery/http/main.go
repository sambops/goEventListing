package main

import (
	//"github.com/goEventListing/API/entity"
	"fmt"
	"log"
	"net/http"

	eventRepo "github.com/goEventListing/API/event/repository"
	eventService "github.com/goEventListing/API/event/services"
	"github.com/goEventListing/API/user/repository"
	"github.com/goEventListing/API/user/services"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/goEventListing/API/delivery/http/handler"

	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:bura@localhost/goeventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	router := httprouter.New()

	// errs := dbconn.CreateTable(entity.Tag{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(err)
	// }

	//user
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/el/user/:id", userHandler.GetUser)

	router.POST("/el/user/login", userHandler.AuthenticateUser)
	router.POST("/el/user/register", userHandler.RegisterUser)
	router.PUT("/el/user/edit", userHandler.EditUser)
	router.POST("/el/user/remove", userHandler.DeleteUser)
	//router.GET("/el/user/logout",userHandler.Logout)

	//dbconn.AutoMigrate(&database.Event{},&database.EventTag{},&database.Tag{},&database.User{},&database.UserTag{})

	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	router.GET("/el/event/allevents", eventHandler.AllEvents)
	router.GET("/el/event/event/:id", eventHandler.Event)
	router.GET("/el/event/upcoming", eventHandler.UpcomingEvents)
	router.POST("/el/event/create", eventHandler.CreateEvent)
	router.GET("/el/event/foru/:id", eventHandler.GetUserSpecificEvent)

	fmt.Printf("...")
	log.Fatal(http.ListenAndServe(":8081", router))

}
