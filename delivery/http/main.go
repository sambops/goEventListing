package main

import (
	"net/http"

	"github.com/goEventListing/delivery/http/handler"
	"github.com/goEventListing/user/services"
	_ "github.com/lib/pq"
	"database/sql"
	"html/template"

	"github.com/goEventListing/user/repository"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/goeventlisting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()


	if err := dbconn.Ping(); err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	userRepo := repository.NewUserRepositoryImpl(dbconn)
	userService := services.NewUserServiceImpl(userRepo)

	userHandler := handler.NewUserHandler(tmpl, userService)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	

	http.HandleFunc("/", userHandler.Index)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/register",userHandler.Register)
	http.HandleFunc("/logout",userHandler.Logout)

	http.ListenAndServe(":8080", nil)

}
		