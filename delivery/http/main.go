package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/birukbelay/Aprojects/goEventListing/delivery/http/handler"
	_ "github.com/birukbelay/Aprojects/goEventListing/entity"
	_ "github.com/lib/pq"

	"github.com/birukbelay/Aprojects/goEventListing/event/repository"
	"github.com/birukbelay/Aprojects/goEventListing/event/services"

	"github.com/birukbelay/Aprojects/goEventListing/review/repository"
	"github.com/birukbelay/Aprojects/goEventListing/review/services"
)

func main() {
	dbconn, err := sql.Open("postgres", "postgres://postgress:bura@localhost/eventListingdb?sslmode = disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("ui/templates/*"))
	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	reviewRepo := repository.NewReviewRepoImpl(dbconn)
	ReviewServ := services.NewReviewServiceImpl(reviewRepo)
	reviewHandler := handler.NewReviewHandler(tmpl, ReviewServ)

	eventRepo := repository.NewEventRepoImp(dbconn)
	eventService := services.NewEventServicesImpl(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	// userRepo := repository.NewUserRepositoryImpl(dbconn)
	// userService := services.NewUserServiceImpl(userRepo)
	// userHandler := handler.NewUserHandler(tmpl, userService)

	// http.HandleFunc("/", userHandler.userIndex)
	// http.HandleFunc("/login", userHandler.Login)
	// http.HandleFunc("/register", userHandler.Register)
	// http.HandleFunc("/logout", userHandler.Logout)

	http.HandleFunc("/", eventHandler.Event)
	fmt.Println("server starting")
	http.ListenAndServe(":8080", nil)

}
