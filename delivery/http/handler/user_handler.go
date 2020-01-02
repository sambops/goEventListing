package handler

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/EventListing/entity"

	uuid "github.com/satori/go.uuid"

	"github.com/EventListing/user"
)

//UserHandler handles user related requests
type UserHandler struct {
	tmpl    *template.Template
	userSrv user.UserService
}

var dbSessions = map[string]string{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(T *template.Template, US user.UserService) *UserHandler {
	return &UserHandler{tmpl: T, userSrv: US}
}

//********************************************************************************************************

//checks whether the user is already logged in or not
func (uh *UserHandler) alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, errr := uh.userSrv.GetUser(un)
	if errr != nil {
		return false
	}
	return true

}
func (uh *UserHandler) getUser(w http.ResponseWriter, req *http.Request) entity.User {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)
	//if the user exists already,get user
	var u entity.User
	if un, ok := dbSessions[c.Value]; ok {
		u, _ = uh.userSrv.GetUser(un)
	}
	return u
}

//*****************************************************************************************************

//Index ... home page
func (uh *UserHandler) Index(w http.ResponseWriter, req *http.Request) {
	u := uh.getUser(w, req)
	uh.tmpl.ExecuteTemplate(w, "home.html", u)
}

//Login handle request on route /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if uh.alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		userName := r.FormValue("userName")
		password := r.FormValue("password")

		_, err := uh.userSrv.AuthenticateUser(userName, password)
		if err != nil {
			panic(err)
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = userName
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return

	}
	uh.tmpl.ExecuteTemplate(w, "login.html", nil)
}

//Register ... handles request on /register
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if uh.alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u entity.User
	if r.Method == http.MethodPost {
		fn := r.FormValue("FirstName")
		ln := r.FormValue("LastName")
		un := r.FormValue("UserName")
		email := r.FormValue("Email")
		pass := r.FormValue("Password")
		phone := r.FormValue("Phone")
		img := r.FormValue("Image")

		_, err := uh.userSrv.GetUser(un)
		if err != nil {
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

		//create a session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		//store user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		//?? what should i put int he place of user id???????????
		u = entity.User{0, fn, ln, un, email, bs, phone, img}

		uh.userSrv.RegisterUser(u)
		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uh.tmpl.ExecuteTemplate(w, "signup.html", u)

}

//Logout ...
func (uh *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	if !uh.alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	//delete the session
	delete(dbSessions, c.Value)
	//remove the cooke
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

}
