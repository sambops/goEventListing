package main

import (
	//"github.com/goEventListing/API/entity"
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
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	router :=httprouter.New()

	// errs := dbconn.CreateTable(entity.Role{}).GetErrors()

	// if len(errs) > 0 {
	// 	panic(err)
	// }

	//user
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService)

	
	
	//router.GET("/el/user/:id", userHandler.GetUser)
	router.GET("/el/user",userHandler.GetUsers)

	//router.POST("/el/user/login", userHandler.)
	router.POST("/el/user/register",userHandler.RegisterUser)
	router.PUT("/el/user/edit",userHandler.EditUser)
	router.POST("/el/user/remove/:id",userHandler.DeleteUser)
	router.POST("/el/user/role/:user",userHandler.UserRoles)

	router.GET("/el/user/check/phone/:phone",userHandler.PhoneExists)
	router.GET("/el/user/check/email/:email",userHandler.EmailExists)
	router.GET("/el/user/username/:username",userHandler.GetUserByUserName)

	//router.GET("/el/user/logout",userHandler.Logout)

	userRoleRepo := repository.NewRoleGormRepo(dbconn)
	userRoleService := services.NewRoleServiceImpl(userRoleRepo)
	userRoleHandler := handler.NewUserRoleHandler(userRoleService)

	router.GET("/el/role/roles",userRoleHandler.Roles)
	router.GET("/el/role/role/:id",userRoleHandler.Role)
	router.GET("/el/role/rolebyname/:name",userRoleHandler.RoleByName)
	router.PUT("/el/role/update",userRoleHandler.UpdateRole)
	router.POST("/el/role/remove/:id",userRoleHandler.DeleteRole)
	router.POST("/el/role/store",userRoleHandler.StoreRole)
	


	userSessionRepo := repository.NewSessionGormRepo(dbconn)
	userSessionService := services.NewSessionService(userSessionRepo)
	userSessionHandler := handler.NewUserSessionHandler(userSessionService)


	router.GET("/el/session/session/:sid",userSessionHandler.Session)
	router.POST("/el/session/store",userSessionHandler.StoreSession)
	router.POST("/el/session/remove/:id",userSessionHandler.DeleteSession)
	
	//dbconn.AutoMigrate(&entity.Event{},&entity.Tag{},&entity.User{},&entity.Review{},&entity.Role{},&entity.Session{})

	//event
	eventRepo := eventRepo.NewEventRepoImp(dbconn)
	eventService := eventService.NewEventServiceImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	router.GET("/el/event/allevents",eventHandler.AllEvents)
	router.GET("/el/event/event/:id",eventHandler.Event)
	router.GET("/el/event/upcoming",eventHandler.UpcomingEvents)
	router.POST("/el/event/create",eventHandler.CreateEvent)
	router.PUT("/el/event/update",eventHandler.UpdateEvent)
	router.GET("/el/event/foru",eventHandler.GetUserSpecificEvent)
	router.POST("/el/event/remove/:id",eventHandler.RemoveEvent)

	//review 
	reviewRepo := rr.NewReviewGormRepo(dbconn)
	reviewservice := rs.NewReviewServiceImpl(reviewRepo)	
	reviewHandler := handler.NewReviewHandler(reviewservice)

	router.GET("/el/reviews", reviewHandler.Reviews)
	//router.GET("/el/review/:id", reviewHandler.Review)
	//router.GET("/el/review/:id", reviewHandler.GetMyReviews)
	router.GET("/el/review/event/:id", reviewHandler.EventReviews)
	router.POST("/el/review/make", reviewHandler.MakeReview)
	router.PUT("/el/review/edit", reviewHandler.EditReview)
	router.POST("/el/review/delete/:id", reviewHandler.DeleteReview)

	http.ListenAndServe(":8181", router)

}
		