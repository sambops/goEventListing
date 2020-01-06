package main

import (
	"net/http"
	"html/template"
)


var tmpl = template.Must(template.ParseGlob("../../ui/templates/*"))

func main() {
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))


	// http.HandleFunc("/", userHandler.Index)
	// http.HandleFunc("user/login", userHandler.Login)
	// http.HandleFunc("user/register",userHandler.Register)
	// http.HandleFunc("user/logout",userHandler.Logout)


	http.ListenAndServe(":8080",nil)
}