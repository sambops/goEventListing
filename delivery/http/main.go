package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/birukbelay/Aprojects/goEventListing/delivery/http/handler"

	"github.com/birukbelay/Aprojects/goEventListing/event/repository"
	"github.com/birukbelay/Aprojects/goEventListing/event/services"

	"github.com/birukbelay/Aprojects/goEventListing/user/repository"
	"github.com/birukbelay/Aprojects/goEventListing/user/services"

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
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	reviewRepo := repository.NewReviewRepositoryImpl(dbconn)
	ReviewServ := services.NewUserServiceImpl(reviewRepo)
	reviewHandler := handler.NewReviewHandler(templ, ReviewServ)

	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)

	userHandler := handler.NewUserHandler(tmpl, userService)


	
	http.HandleFunc("/", userHandler.userIndex)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/logout", userHandler.Logout)

	fmt.Println("server starting")
	http.ListenAndServe(":8080", nil)

}
