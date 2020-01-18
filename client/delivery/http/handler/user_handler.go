package handler

import (
<<<<<<< HEAD
	"fmt"

	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"github.com/julienschmidt/httprouter"

	"html/template"

	"golang.org/x/crypto/bcrypt"

	"net/http"

	uuid "github.com/satori/go.uuid"
=======
	"github.com/julienschmidt/httprouter"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"

	"html/template"


	"golang.org/x/crypto/bcrypt"
	
	"net/http"

	uuid "github.com/satori/go.uuid"

>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
)

//UserHandler handles user related requests
type UserHandler struct {
<<<<<<< HEAD
	tmpl *template.Template
}

=======
	tmpl   *template.Template

}


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
var dbSessions = map[string]uint{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(T *template.Template) *UserHandler {
	return &UserHandler{tmpl: T}
}

<<<<<<< HEAD
//checks whether the user is already logged in or not
func alreadyLoggedIn(req *http.Request) bool {
=======
//AlreadyLoggedIn .... checks whether the user is already logged in or not
func AlreadyLoggedIn(req *http.Request) bool {
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	id := dbSessions[c.Value]
	_, errr := service.GetUser(id)
<<<<<<< HEAD

=======
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if errr != nil {
		return false
	}
	return true

}
<<<<<<< HEAD
func getUser(w http.ResponseWriter, req *http.Request) *entity.User {
=======
//GetUser .... gets currently logged user
func GetUser(w http.ResponseWriter, req *http.Request) *entity.User {
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
<<<<<<< HEAD
=======
			MaxAge: 60 * 3,
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		}
	}
	http.SetCookie(w, c)

	//if the user exists already,get user
	var u *entity.User
	if id, ok := dbSessions[c.Value]; ok {
		u, _ = service.GetUser(id)
	}
	return u
<<<<<<< HEAD

}

//Index ... home page before login
func (uh *UserHandler) Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	u := getUser(w, req)
	fmt.Printf("Hellow")
=======
	
}


//Index ... home page before login
func (uh *UserHandler) Index(w http.ResponseWriter, req *http.Request,ps httprouter.Params) {
	u := GetUser(w, req)
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	uh.tmpl.ExecuteTemplate(w, "home.html", u)
}

//Login handle request on route /login
<<<<<<< HEAD
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if alreadyLoggedIn(r) {
=======
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	if AlreadyLoggedIn(r) {
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		userName := r.FormValue("uname")
		password := r.FormValue("psw")

		usr, err := service.AuthenticateUser(userName, password)
		if err != nil {
			//panic(err)
<<<<<<< HEAD
			http.Error(w, "hey check what u wrote please", 404)
=======
			http.Error(w,"hey check what u wrote please",404)
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = usr.ID
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return

	}
	uh.tmpl.ExecuteTemplate(w, "login.html", nil)
}

<<<<<<< HEAD
//Register ... handles request on /register
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if alreadyLoggedIn(r) {
=======

//Register ... handles request on /register
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	if AlreadyLoggedIn(r) {
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u *entity.User
	if r.Method == http.MethodPost {
		fn := r.FormValue("FirstName")
		ln := r.FormValue("LastName")
		un := r.FormValue("UserName")
		email := r.FormValue("Email")
		pass := r.FormValue("Password")
		phone := r.FormValue("Phone")
		img := r.FormValue("Image")

		_, err := service.GetUserByUserName(un)
		if err != nil {
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

<<<<<<< HEAD
=======
		
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		//store user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		//?? what should i put int he place of user id???????????
<<<<<<< HEAD
		u = &entity.User{FirstName: fn, LastName: ln, UserName: un, Email: email, Password: bs, Phone: phone, Image: img}

=======
		u = &entity.User{FirstName:fn,LastName:ln,UserName:un,Email:email,Password:bs,Phone:phone,Image:img}
		
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		//create a session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = u.ID

		service.RegisterUser(u)
		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uh.tmpl.ExecuteTemplate(w, "signup.html", u)

}

//Logout ...
<<<<<<< HEAD
func (uh *UserHandler) Logout(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if !alreadyLoggedIn(req) {
=======
func (uh *UserHandler) Logout(w http.ResponseWriter, req *http.Request,ps httprouter.Params) {
	if !AlreadyLoggedIn(req) {
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
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
<<<<<<< HEAD
	http.Redirect(w, req, "/", http.StatusSeeOther)
=======
	http.Redirect(w,req,"/", http.StatusSeeOther)
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
}
