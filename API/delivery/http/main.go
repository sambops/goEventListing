package main

import (
<<<<<<< HEAD
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
=======
	"github.com/goEventListing/API/user/services"
	"github.com/goEventListing/API/user/repository"
	"github.com/julienschmidt/httprouter"
	eventRepo"github.com/goEventListing/API/event/repository"
	eventService "github.com/goEventListing/API/event/services"
	
	rr "github.com/goEventListing/API/review/repository"
	rs "github.com/goEventListing/API/review/services"
	
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/goEventListing/API/delivery/http/handler"
	
	_ "github.com/lib/pq"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/goeventlisting?sslmode=disable")
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

<<<<<<< HEAD
	router := httprouter.New()
	dbconn.AutoMigrate(&entity.Event{}, &entity.EventTag{}, &entity.Tag{}, &entity.User{}, &entity.Review{})
	// createTables(dbconn)

	//review handler
	reviewRepo := rr.NewReviewGormRepo(dbconn)
	reviewservice := rs.NewReviewServiceImpl(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewservice)

	router.GET("/el/reviews", reviewHandler.Reviews)
	router.GET("/el/review/single/review/:ids", reviewHandler.Review)

	router.GET("/el/user/review/:id", reviewHandler.GetMyReviews)

	router.GET("/el/event/reviews/:id", reviewHandler.EventReviews)

	router.POST("/el/review/make", reviewHandler.MakeReview)

	router.PUT("/el/review/edit/", reviewHandler.EditReview)
	router.DELETE("/el/review/delete/:id", reviewHandler.DeleteReview)

	// user hanndler
=======

	router :=httprouter.New()

	

	// errs := dbconn.CreateTable(entity.Tag{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(err)
	// }


	//user
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)

<<<<<<< HEAD
	router.GET("/el/user/my/:id", userHandler.GetUser)

	router.POST("/el/user/login", userHandler.AuthenticateUser)
	router.POST("/el/user/register", userHandler.RegisterUser)
	router.PUT("/el/user/edit", userHandler.EditUser)
	router.POST("/el/user/remove", userHandler.DeleteUser)
	//router.GET("/el/user/logout",userHandler.Logout)

	//event handler
=======
	
	
	router.GET("/el/user/:id", userHandler.GetUser)

	router.POST("/el/user/login", userHandler.AuthenticateUser)
	router.POST("/el/user/register",userHandler.RegisterUser)
	router.PUT("/el/user/edit",userHandler.EditUser)
	router.POST("/el/user/remove",userHandler.DeleteUser)
	//router.GET("/el/user/logout",userHandler.Logout)
	
	//dbconn.AutoMigrate(&database.Event{},&database.Tag{},&database.User{},&entity.Review{})

	//event
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

<<<<<<< HEAD
	router.GET("/el/event/allevents", eventHandler.AllEvents)
	router.GET("/el/event/event/:id", eventHandler.Event)
	router.GET("/el/event/upcoming", eventHandler.UpcomingEvents)
	router.POST("/el/event/create", eventHandler.CreateEvent)
	router.GET("/el/event/foru/:id", eventHandler.GetUserSpecificEvent)

	fmt.Printf("...8081...")
	log.Fatal(http.ListenAndServe(":8081", router))

}
=======
	router.GET("/el/event/allevents",eventHandler.AllEvents)
	router.GET("/el/event/event/:id",eventHandler.Event)
	router.GET("/el/event/upcoming",eventHandler.UpcomingEvents)
	router.POST("/el/event/create",eventHandler.CreateEvent)
	router.PUT("/el/event/update",eventHandler.UpdateEvent)
	router.GET("/el/event/foru/:id",eventHandler.GetUserSpecificEvent)
	router.POST("/el/event/remove",eventHandler.RemoveEvent)

	//review 
	reviewRepo := rr.NewReviewGormRepo(dbconn)
	reviewservice := rs.NewReviewServiceImpl(reviewRepo)	
	reviewHandler := handler.NewReviewHandler(reviewservice)

	router.GET("/el/reviews", reviewHandler.Reviews)
	router.GET("/el/review/:id", reviewHandler.Review)
	// router.GET("/el/review/:id", reviewHandler.GetMyReviews)
	// router.GET("/el/review/:id", reviewHandler.EventReviews)
	router.POST("/el/review/make", reviewHandler.MakeReview)
	router.PUT("/el/review/edit/", reviewHandler.EditReview)
	router.DELETE("/el/review/delete/:id", reviewHandler.DeleteReview)

	http.ListenAndServe(":8181", router)

}
		
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
