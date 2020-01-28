package main

import (
	"github.com/goEventListing/client/rtoken"
	"time"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/delivery/http/handler"
	"net/http"
	"html/template"
)



func main() {
		
	tmpl := template.Must(template.ParseGlob("../../ui/templates/*.html"))
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))


	sess := configSess()


	// http.HandleFunc("/", userHandler.Index)
	// http.HandleFunc("user/login", userHandler.Login)
	// http.HandleFunc("user/register",userHandler.Register)
	// http.HandleFunc("user/logout",userHandler.Logout)

	//user
	usrHandler := handler.NewUserHandler(tmpl,sess,csrfSignKey)
	
	//router.GET("/",usrHandler.Index)
	http.HandleFunc("/el/user/register",usrHandler.Register)
	http.HandleFunc("/el/user/login",usrHandler.Login)
	http.HandleFunc("/el/user/logout",usrHandler.Logout)
	http.HandleFunc("/",usrHandler.CheckIndex)
	

	//event
	// evtHandler := handler.NewEventHandler(tmpl)
	// router.GET("/el/event/all",evtHandler.Events)
	// router.GET("/el/event/upcoming",evtHandler.Upcoming)
	// router.GET("/el/event/create",evtHandler.CreateEvent)
	// router.GET("/el/event/foru",evtHandler.UserSpecific)


	http.ListenAndServe(":8080",nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
