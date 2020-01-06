package main

import (
	"github.com/julienschmidt/httprouter"
	eventRepo"github.com/goEventListing/API/event/repository"
	eventService "github.com/goEventListing/API/event/services"
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/goEventListing/API/delivery/http/handler"
	
	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/goeventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()


	router :=httprouter.New()
	//tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
		
	//fs := http.FileServer(http.Dir("ui/assets"))
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//user
	// userRepo := repository.NewUserRepositoryImpl(dbconn)
	// userService := services.NewUserServiceImpl(userRepo)
	// userHandler := handler.NewUserHandler(tmpl, userService)

	
	
	// router.GET("/", userHandler.Index)
	// router.POST("/el/user/login", userHandler.Login)
	// router.POST("/el/user/register",userHandler.Register)
	// router.PUT("/el/user/edit",userHandler.Register)
	// router.POST("/el/user/remove",userHandler.Register)
	// router.GET("/el/user/logout",userHandler.Logout)
	//dbconn.AutoMigrate(&database.Event{},&database.EventTag{},&database.Tag{},&database.User{},&database.UserTag{})

	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	router.GET("/el/event/allevents",eventHandler.AllEvents)
	router.GET("/el/event/event/:id",eventHandler.Event)
	router.GET("/el/event/upcoming",eventHandler.UpcomingEvents)
	router.POST("/el/event/create",eventHandler.CreateEvent)
	router.GET("/el/event/foru",eventHandler.GetUserSpecificEvent)

	
	

	





	http.ListenAndServe(":8080", nil)

}
		