package handler

import (
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"

	"html/template"


	"golang.org/x/crypto/bcrypt"
	
	"net/http"

	uuid "github.com/satori/go.uuid"

)

//UserHandler handles user related requests
type UserHandler struct {
	tmpl   *template.Template
}


var dbSessions = map[string]uint{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(T *template.Template) *UserHandler {
	return &UserHandler{tmpl: T}
}

//checks whether the user is already logged in or not
func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	id := dbSessions[c.Value]
	_, errr := service.GetUser(id)
	
	if errr != nil {
		return false
	}
	return true

}
func getUser(w http.ResponseWriter, req *http.Request) *entity.User {
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
	var u *entity.User
	if id, ok := dbSessions[c.Value]; ok {
		u, _ = service.GetUser(id)
	}
	return u

	
}


//Index ... home page before login
func (uh *UserHandler) Index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	uh.tmpl.ExecuteTemplate(w, "home.html", u)
}

//Login handle request on route /login
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		userName := r.FormValue("uname")
		password := r.FormValue("psw")

		usr, err := service.AuthenticateUser(userName, password)
		if err != nil {
			//panic(err)
			http.Error(w,"hey check what u wrote please",404)
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


//Register ... handles request on /register
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
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

		
		//store user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		//?? what should i put int he place of user id???????????
		u = &entity.User{FirstName:fn,LastName:ln,UserName:un,Email:email,Password:bs,Phone:phone,Image:img}
		
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
func (uh *UserHandler) Logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
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
	http.Redirect(w,req,"/", http.StatusSeeOther)
}
