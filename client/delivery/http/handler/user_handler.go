package handler

import (
	"fmt"
	"github.com/goEventListing/client/permission"
	"github.com/goEventListing/client/form"
	"github.com/goEventListing/client/rtoken"
	"github.com/goEventListing/client/session"
	"context"
	"net/url"
	"strings"
	"strconv"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	
)

//UserHandler handles user related requests
type UserHandler struct {
	tmpl   *template.Template
	//sessionService
	userSess       *entity.Session
	loggedInUser   *entity.User
	//userRoleService
	csrfSignKey    []byte
}

type contextKey string
var ctxUserSessionKey = contextKey("signed_in_user_session")

//var dbSessions = map[string]uint{} //session ID,user ID

//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(T *template.Template,usrSess *entity.Session, csKey []byte) *UserHandler {
	return &UserHandler{tmpl: T,userSess: usrSess, csrfSignKey: csKey}
}

// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorized checks if a user has proper authority to access a give route
func (uh *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		roles,err := service.UserRoles(uh.loggedInUser)
		//roles, errs := uh.userService.UserRoles(uh.loggedInUser)
		if  err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}


// Register hanldes the GET/POST /signup requests
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register")
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values 
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signup.layout", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		singnUpForm.Required("fullName", "email", "password")
		singnUpForm.MatchesPattern("email", form.EmailRX)
		singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		//singnUpForm.PasswordMatches("password", "confirmpassword")
		singnUpForm.CSRF = token

		//If there are any errors, redisplay the signup form.
		// if !singnUpForm.Valid() {
		// 	fmt.Println("err1")
		// 	fmt.Println(err)
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

		// pExists := service.PhoneExists(r.FormValue("phone"))
		// if *pExists {
		// 	fmt.Println("err2")

		// 	singnUpForm.VErrors.Add("phone", "Phone Already Exists")
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }
		// eExists := service.EmailExists(r.FormValue("email"))
		// if *eExists {
		// 	fmt.Println("err3")
		// 	singnUpForm.VErrors.Add("email", "Email Already Exists")
		// 	uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
		// 	return
		// }

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("psw")), 12)
		if err != nil {
			fmt.Println("err4")
			singnUpForm.VErrors.Add("psw", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		role,err := service.RoleByName("USER")
		//role, errs := uh.userRole.RoleByName("USER")

		if err != nil {
			fmt.Println("err5")
			fmt.Println(err)
			singnUpForm.VErrors.Add("role", "could not assign role to the user")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		user := &entity.User{
			UserName: r.FormValue("userName"),
			FirstName:r.FormValue("firstName"),
			LastName:r.FormValue("lastName"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: hashedPassword,
			RoleID:   role.ID,
		}
		_,err = service.RegisterUser(user)
		//_, errs = uh.userService.StoreUser(user)
		if err!= nil {
			fmt.Println("err6")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/el/user/login", http.StatusSeeOther)
	}
}
// Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(uh.userSess.UUID, w)
	service.DeleteSession(uh.userSess.UUID)
	uh.loggedInUser = nil
	//uh.sessionService.DeleteSession(userSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Login hanldes the GET/POST /login requests
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "signin.layout", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			fmt.Println("err1")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		
		usr,err := service.GetUserByUserName(r.FormValue("userName"))
		if err != nil {
			fmt.Println("err 2")
			loginForm.VErrors.Add("generic", "UserName or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "signin.layout", loginForm)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(r.FormValue("psw")))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loginForm.VErrors.Add("generic", "Your email address or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "signin.layout", loginForm)
			return
		}

		uh.loggedInUser = usr
		fmt.Println(uh.loggedInUser)
		claims := rtoken.Claims(usr.Email, uh.userSess.Expires)
		session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
		newSess,err := service.StoreSession(uh.userSess)
		if err != nil {
			fmt.Println("err 3")
			loginForm.VErrors.Add("generic", "Failed to store session")
			uh.tmpl.ExecuteTemplate(w, "signin.layout", loginForm)
			return
		}
		uh.userSess = newSess
		roles,_ := service.UserRoles(usr)

		if uh.checkAdmin(roles) {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		// uh.tmpl.ExecuteTemplate(w,"all.layout",loginForm)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}


func (uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess 
	c, err := r.Cookie(userSess.UUID)
	if err != nil {
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		return false
	}
	return true
}


// AdminUsers handles Get /admin/users request
func (uh *UserHandler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	users,err := service.GetUsers()
	//users, errs := uh.userService.Users()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Users   []entity.User
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Users:   *users,
		CSRF:    token,
	}
	uh.tmpl.ExecuteTemplate(w, "admin.users.layout", tmplData)
}


// AdminUsersNew handles GET/POST /admin/users/new request
func (uh *UserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		roles,err := service.Roles()
		//roles, errs := uh.userRole.Roles()
		if err!= nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		accountForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   *[]entity.Role
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			Roles:   roles,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate the form contents
		accountForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		accountForm.Required("fullname", "email", "password", "confirmpassword")
		accountForm.MatchesPattern("email", form.EmailRX)
		accountForm.MatchesPattern("phone", form.PhoneRX)
		accountForm.MinLength("password", 8)
		accountForm.PasswordMatches("password", "confirmpassword")
		accountForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !accountForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		pExists := service.PhoneExists(r.FormValue("phone"))
		//pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		if *pExists {
			accountForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		//eExists := uh.userService.EmailExists(r.FormValue("email"))
		eExists := service.EmailExists(r.FormValue("email"))

		if *eExists {
			accountForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			accountForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			accountForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		user := &entity.User{
			FirstName :r.FormValue("firstName"),
			LastName: r.FormValue("lastName"),
			UserName: r.FormValue("userName"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: hashedPassword,
			RoleID:   uint(roleID),
		}
		_, err = service.RegisterUser(user)

		//_, errs := uh.userService.StoreUser(user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}


func (uh *UserHandler) checkAdmin(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}


// AdminUsersUpdate handles GET/POST /admin/users/update?id={id} request
func (uh *UserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		
		id, err := strconv.Atoi(idRaw) 
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		//user, errs := uh.userService.User(uint(id))
		user,err := service.GetUser(uint(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		//roles, errs := uh.userRole.Roles()
		roles,err := service.Roles()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		//role, errs := uh.userRole.Role(user.RoleID)
		role,err := service.Role(user.RoleID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values := url.Values{}
		values.Add("userid", idRaw)
		values.Add("fullname", user.UserName)
		values.Add("email", user.Email)
		values.Add("role", string(user.RoleID))
		values.Add("phone", user.Phone)
		values.Add("rolename", role.Name)

		upAccForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Roles   *[]entity.Role
			User    *entity.User
			CSRF    string
		}{
			Values:  values,
			VErrors: form.ValidationErrors{},
			Roles:   roles,
			User:    user,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		upAccForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		upAccForm.Required("fullname", "email", "phone")
		upAccForm.MatchesPattern("email", form.EmailRX)
		upAccForm.MatchesPattern("phone", form.PhoneRX)
		upAccForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !upAccForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		userID := r.FormValue("userid")
		uid, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		//user, errs := uh.userService.User(uint(uid))
		user,err := service.GetUser(uint(uid))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		eExists := service.EmailExists(r.FormValue("email"))
		//eExists := uh.userService.EmailExists(r.FormValue("email"))
		if (user.Email != r.FormValue("email")) && *eExists {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		pExists := service.PhoneExists(r.FormValue("phone"))
		//pExists := uh.userService.PhoneExists(r.FormValue("phone"))

		if (user.Phone != r.FormValue("phone")) && *pExists {
			upAccForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			uh.tmpl.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
			return
		}
		usr := &entity.User{
			ID:       user.ID,
			UserName: r.FormValue("userName"),
			FirstName: r.FormValue("firstName"),
			LastName:r.FormValue("lastName"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: user.Password,
			RoleID:   uint(roleID),
		}
		//_, errs = uh.userService.UpdateUser(usr)
		_,err = service.EditUser(usr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersDelete handles Delete /admin/users/delete?id={id} request
func (uh *UserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		_,err = service.DeleteUser(uint(id))
		//_, errs := uh.userService.DeleteUser(uint(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

// CheckIndex ...
func (uh *UserHandler) CheckIndex(w http.ResponseWriter, r *http.Request) {
	data, err := service.AllEvents()
	d := struct {
		Data *[]entity.Event
		LoggedIn bool
	}{
		Data: data,
		LoggedIn: uh.loggedInUser != nil,
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uh.tmpl.ExecuteTemplate(w, "all.layout", d)
	return
}

//Events handle reques on route/events
func(uh *UserHandler) Events(w http.ResponseWriter,req *http.Request){
	
	evt,err := service.AllEvents()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Events 		*[]entity.Event
		CSRF       string
		LoggedIn bool
	}{
		Values:     nil,
		VErrors:    nil,
		Events: 	evt,
		CSRF:       token,
		LoggedIn: uh.loggedInUser != nil,
	}

	//fmt.Println("events:",evt)

	
	uh.tmpl.ExecuteTemplate(w, "all.layout", tmplData)

}


//Upcoming handle request on route/upcoming
func(uh *UserHandler) Upcoming(w http.ResponseWriter,req *http.Request){
	
	upcoming,err := service.UpcomingEvent()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Upcoming *[]entity.Event
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Upcoming: 	upcoming,
		CSRF:       token,
		
	}
	fmt.Println("upcoming:", upcoming)
	uh.tmpl.ExecuteTemplate(w, "thisweekend.layout", tmplData)
}

//CreateEvent ... request on route/create
func(uh *UserHandler) CreateEvent(w http.ResponseWriter,req *http.Request){

	if uh.loggedInUser == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
   
   token, err := rtoken.CSRFToken(uh.csrfSignKey)
   if err != nil {
	   fmt.Println("stack here")
	   http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
   }
   if req.Method == http.MethodGet {
	   newCatForm := struct {
		   Values  url.Values
		   VErrors form.ValidationErrors
		   CSRF    string
	   }{
		   Values:  nil,
		   VErrors: nil,
		   CSRF:    token,
	   }
	   uh.tmpl.ExecuteTemplate(w, "addEvent.layout", newCatForm)
   }
   if req.Method == http.MethodPost {
	   // Parse the form data
	   err := req.ParseForm()
	   if err != nil {
		   http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		   return
	   }
	   // Validate the form contents
	//    newEvtForm := form.Input{Values: req.PostForm, VErrors: form.ValidationErrors{}}
	//    newEvtForm.Required("name", "details","country","city","place","price")
	//    newEvtForm.MinLength("details", 15)
	//    newEvtForm.CSRF = token
	//    // If there are any errors, redisplay the signup form.
	//    if !newEvtForm.Valid() {
	// 	   fmt.Println("i'm here")
	// 	   uh.tmpl.ExecuteTemplate(w, "addEvent.html", newEvtForm)
	// 	   return
	//    }
	   
	   // mf, fh, err := req.FormFile("image")
	   // if err != nil {
	   // 	newEvtForm.VErrors.Add("image", "File error")
	   // 	eh.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newEvtForm)
	   // 	return
	   // }
	   //defer mf.Close()
	  prc := req.FormValue("price")
	  fmt.Println("price: " + prc)
	   price,err := strconv.ParseFloat(prc, 64)
	   if err != nil {
		   panic(err)
	   }
	   evt := &entity.Event{
		   Name:     req.FormValue("name"),
		   Details : req.FormValue("details"),
		   Country : req.FormValue("country"),
		   City : req.FormValue("city"),
		   Place : req.FormValue("place"),
		   Price : &price,
		   UserID : uh.loggedInUser.ID,
		   Image:  "img.jpg",
		   //fh.Filename,
	   }
	   fmt.Println(evt)
	   //writeFile(&mf, fh.Filename)
	   //_, errs := ach.categorySrv.StoreCategory(ctg)
	   _,err = service.AddEvent(evt)
	   if err != nil {
		   fmt.Println("last err")
		   http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	   }
	   fmt.Println("finally reach here")
	   http.Redirect(w, req, "/", http.StatusSeeOther)
   }
}

//UserSpecific handle request on route/upcoming
func (uh *UserHandler) UserSpecific(w http.ResponseWriter,req *http.Request){
	
	// eh.tmpl.ExecuteTemplate(w,"foru.html",tmplData)
	evnts,err := service.GetUserSubscribedEvents(uh.loggedInUser.ID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		UserSpecific *[]entity.Event
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		UserSpecific: evnts,
		CSRF:       token,
	}
	fmt.Println("user specific:",evnts)
	uh.tmpl.ExecuteTemplate(w, "foru.html", tmplData)
}


//RemoveEvent ... handle request on route/remove
func (uh *UserHandler) RemoveEvent(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_,err = service.DeleteEvent(uint (id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}



//REVIEW

//EventReviews ....
func (uh *UserHandler) EventReviews(w http.ResponseWriter,req *http.Request){                                                                                        
	
	eventid,_ := strconv.Atoi(req.URL.Query().Get("id"))
	fmt.Println(eventid)
	review,err := service.EventReviews (uint(eventid))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		reviews *[]entity.Review
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		reviews : review,
		CSRF:       token,
	}
	fmt.Println("reviews:",review)
	uh.tmpl.ExecuteTemplate(w, "create.html", tmplData)

}

//MakeReviewAndRating ... handlers request on /el/review/make
func (uh *UserHandler) MakeReviewAndRating(w http.ResponseWriter,req *http.Request){
	id := uh.loggedInUser.ID

	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if req.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.tmpl.ExecuteTemplate(w, "newEvent.html", newCatForm)
	}else if req.Method == http.MethodPost {
			// Parse the form data
			err := req.ParseForm()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// Validate the form contents
			newRvwForm := form.Input{Values: req.PostForm, VErrors: form.ValidationErrors{}}
			//newRvwForm.Required("name", "details","country","city","place","price")
			newRvwForm.MinLength("message", 2)
			newRvwForm.CSRF = token
			// If there are any errors, redisplay the signup form.
			if !newRvwForm.Valid() {
				fmt.Println("error")
				uh.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newRvwForm)
				return
			}
			rating,err  := strconv.Atoi(req.FormValue("rating"))
			if err != nil {
				panic(err)
			}
			fmt.Println(rating)
			eventid,_ := strconv.Atoi(req.URL.Query().Get("id"))
			message := req.FormValue("message")
		
	
			rvw := &entity.Review{
				Rating: rating,
				UserID : id,
				EventID :uint (eventid),
				Message : message,
			}

			//writeFile(&mf, fh.Filename)
			//_, errs := ach.categorySrv.StoreCategory(ctg)
			_,err = service.MakeReviewAndRating(rvw)
			if err != nil {
				fmt.Println("lemin lemin")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
			}
}

// // UpdateReview ... handles request on /el/review/edit
// func (uh *UserHandler) UpdateReview(w http.ResponseWriter,r *http.Request){
// 		token, err := rtoken.CSRFToken(uh.csrfSignKey)
// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		}
// 		if r.Method == http.MethodGet {
// 			idRaw := r.URL.Query().Get("id")
// 			id, err := strconv.Atoi(idRaw)
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 			}
// 			rvw,err := service.EventReviews(uint (id))
			
// 			//cat, errs := rh.categorySrv.Category(uint(id))
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			}
			
// 			values := url.Values{}
// 			values.Add("catid", idRaw)
// 			values.Add("catname",rvw.)
// 			values.Add("catdesc", cat.Description)
// 			values.Add("catimg", cat.Image)
// 			upCatForm := struct {
// 				Values   url.Values
// 				VErrors  form.ValidationErrors
// 				Category *entity.Category
// 				CSRF     string
// 			}{
// 				Values:   values,
// 				VErrors:  form.ValidationErrors{},
// 				Category: cat,
// 				CSRF:     token,
// 			}
// 			ach.tmpl.ExecuteTemplate(w, "reviewUpdate.html", upCatForm)
// 			return
// 		}
// 		if r.Method == http.MethodPost {
// 			// Parse the form data
// 			err := r.ParseForm()
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 				return
// 			}
// 			// Validate the form contents
// 			updateRvwForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
// 			updateRvwForm.MinLength("message", 5)
// 			updateRvwForm.CSRF = token

// 			rvwID, err := strconv.Atoi(r.FormValue("id"))
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 			}
// 			rvww := &entity.Review{
// 				ID:          uint (rvwID),
// 				UserID : id,
// 				EventID:id,
// 				Message:r.FormValue("message"),
// 			}
// 			// mf, fh, err := r.FormFile("catimg")
// 			// if err == nil {
// 			// 	ctg.Image = fh.Filename
// 			// 	err = writeFile(&mf, ctg.Image)
// 			// }
// 			// if mf != nil {
// 			// 	defer mf.Close()
// 			// }
// 			//_, errs := ach.categorySrv.UpdateCategory(ctg)
// 			_,err = service.UpdateReview(rvww)
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				return
// 			}
// 			http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
// 			return
// 		}


// }


//DeleteReview ... handles request on /el/review/delete
func (uh *UserHandler) DeleteReview(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_,err = service.DeleteReview(uint (id))
		//_, errs := ach.categorySrv.DeleteCategory(uint(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (uh *UserHandler) Event(w http.ResponseWriter, r *http.Request){
	fmt.Println("here")
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		event, err := service.GetEvent(uint(id))
		if err != nil {
			fmt.Println("error 1")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		reviews, err := service.EventReviews(uint(id))
		if err != nil {
			fmt.Println("error 2")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		data := struct {
			Event entity.Event
			Reviews []entity.Review
			LoggedIn bool
		}{
			Event: *event,
			Reviews: *reviews,
			LoggedIn: uh.loggedInUser != nil,
		}

		
		fmt.Println(uh.loggedInUser != nil)
		fmt.Println(event)
		fmt.Println(reviews)

		uh.tmpl.ExecuteTemplate(w, "event.layout", data)
		return
	}
}