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

	fs := http.FileServer(http.Dir("../../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))


	sess := configSess()

	
	

	//user related
	usrHandler := handler.NewUserHandler(tmpl,sess,csrfSignKey)
	
	//router.GET("/",usrHandler.Index)
	http.HandleFunc("/el/user/register",usrHandler.Register)
	http.HandleFunc("/el/user/login",usrHandler.Login)
	http.HandleFunc("/el/user/logout",usrHandler.Logout)
	http.HandleFunc("/",usrHandler.CheckIndex)

	http.HandleFunc("/el/event/all",usrHandler.Events)
	http.HandleFunc("/e1/event", usrHandler.Event)
	http.HandleFunc("/el/event/upcoming",usrHandler.Upcoming)
	http.HandleFunc("/el/event/create",usrHandler.CreateEvent)
	http.HandleFunc("/el/event/foru",usrHandler.UserSpecific)
	http.HandleFunc("/el/event/remove",usrHandler.RemoveEvent)

	http.HandleFunc("/el/review/event",usrHandler.EventReviews)
	http.HandleFunc("/el/review/make",usrHandler.MakeReviewAndRating)
	//http.HandleFunc("/el/review/update",usrHandler.UpdateReview)
	http.HandleFunc("/el/review/delete",usrHandler.DeleteReview)

	//Admin
	http.Handle("/admin/users", usrHandler.Authenticated(usrHandler.Authorized(http.HandlerFunc(usrHandler.AdminUsers))))
	http.Handle("/admin/users/new",usrHandler.Authenticated(usrHandler.Authorized(http.HandlerFunc(usrHandler.AdminUsersNew))))
	http.Handle("/admin/update",usrHandler.Authenticated(usrHandler.Authorized(http.HandlerFunc(usrHandler.AdminUsersUpdate))))
	http.Handle("/admin/delete",usrHandler.Authenticated(usrHandler.Authorized(http.HandlerFunc(usrHandler.AdminUsersDelete))))



	

	//event
	//  evtHandler := handler.NewEventHandler(tmpl,csrfSignKey)

	// http.HandleFunc("/el/event/all",evtHandler.Events)
	// http.HandleFunc("/el/event/upcoming",evtHandler.Upcoming)
	// http.HandleFunc("/el/event/create",evtHandler.CreateEvent)
	// http.HandleFunc("/el/event/foru",evtHandler.UserSpecific)
	// http.HandleFunc("/el/event/remove",evtHandler.RemoveEvent)



	//review
	//rvwHandler := handler.NewReviewHandler(tmpl,csrfSignKey)
	// http.HandleFunc("/el/review/all",rvwHandler.EventReviews)
	// http.HandleFunc("/el/review/make",rvwHandler.MakeReviewAndRating)
	// http.HandleFunc("/el/review/update",rvwHandler.UpdateReview)
	// http.HandleFunc("/el/review/delete",rvwHandler.DeleteReview)





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
