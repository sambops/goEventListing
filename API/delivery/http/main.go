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

	//dbconn.AutoMigrate(&database.Event{},&database.Tag{},&database.User{},&entity.Review{})

	//review handler
	reviewRepo := rr.NewReviewGormRepo(dbconn)
	reviewservice := rs.NewReviewServiceImpl(reviewRepo)
	rH := handler.NewReviewHandler(reviewservice)

	router.GET("/el/reviews", rH.Reviews)
	router.GET("/el/review/single/review/:ids", rH.Review)
	router.GET("/el/event/reviews/:id", rH.EventReviews) //reviews for an event
	router.GET("/el/review/user/:id", rH.GetMyReviews)   //user reviews
	router.POST("/el/review/make", rH.MakeReview)
	router.PUT("/el/review/edit/", rH.EditReview)
	router.DELETE("/el/review/delete/:id", rH.DeleteReview)
	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService, reviewservice)

	router.GET("/el/event/allevents", eventHandler.AllEvents)
	router.GET("/el/event/event/:id", eventHandler.Event)
	router.GET("/el/event/upcoming", eventHandler.UpcomingEvents)
	router.POST("/el/event/create", eventHandler.CreateEvent)
	router.GET("/el/event/foru/:id", eventHandler.GetUserSpecificEvent)
	router.POST("/el/event/remove", eventHandler.RemoveEvent)
	router.PUT("/el/event/update", eventHandler.UpdateEvent)

	fmt.Printf("...8081...")
	log.Fatal(http.ListenAndServe(":8081", router))

}
