package main

import (
	"net/http"

	"github.com/EventListing/delivery/http/handler"
	"github.com/EventListing/user/services"

	"database/sql"
	"html/template"

	"github.com/EventListing/user/repository"
)

func main() {
	dbconn, err := sql.Open("postgres", "postgres://app_admin:post.dalvic.gres@localhost/eventListingdb?sslmode = disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	if err := dbconn.Ping(); err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseGlob("ui/templates/*"))
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
		