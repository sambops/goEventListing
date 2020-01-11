package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goEventListing/API/entity"

	eventRepo "github.com/goEventListing/API/event/repository"
	eventService "github.com/goEventListing/API/event/services"
	"github.com/goEventListing/API/user/repository"
	"github.com/goEventListing/API/user/services"

	rr "github.com/goEventListing/API/review/repository"
	rs "github.com/goEventListing/API/review/services"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	"github.com/goEventListing/API/delivery/http/handler"

	_ "github.com/lib/pq"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Event{}, &entity.Review{}, &entity.User{}, &entity.Notification{}, &entity.Tag{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:bura@localhost/eventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	router := httprouter.New()
	dbconn.AutoMigrate(&entity.Event{}, &entity.EventTag{}, &entity.Tag{}, &entity.User{}, &entity.Review{})
	// createTables(dbconn)

	//review handler
	reviewRepo := rr.NewReviewGormRepo(dbconn)
	reviewservice := rs.NewReviewServiceImpl(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewservice)

	router.GET("/reviews", reviewHandler.Reviews)
	router.GET("/event/review/:id", reviewHandler.Review)

	router.POST("/event/make", reviewHandler.MakeReview)
	router.GET("/user/review/:id", reviewHandler.GetMyReviews)
	router.GET("/event/reviews/:id", reviewHandler.EventReviews)
	router.PUT("/event/review/:id", reviewHandler.PutReview)
	router.GET("/event/revie/:id", reviewHandler.DeleteReview)

	// user hanndler
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/el/user/:id", userHandler.GetUser)

	router.POST("/el/user/login", userHandler.AuthenticateUser)
	router.POST("/el/user/register", userHandler.RegisterUser)
	router.PUT("/el/user/edit", userHandler.EditUser)
	router.POST("/el/user/remove", userHandler.DeleteUser)
	//router.GET("/el/user/logout",userHandler.Logout)

	//event handler
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
