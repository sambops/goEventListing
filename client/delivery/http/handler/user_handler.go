package handler

import (
	"github.com/goEventListing/client/permission"
	"github.com/julienschmidt/httprouter"
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

		for _, role := range *roles {
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
func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
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
		singnUpForm.Required("fullname", "email", "password", "confirmpassword")
		singnUpForm.MatchesPattern("email", form.EmailRX)
		singnUpForm.MatchesPattern("phone", form.PhoneRX)
		singnUpForm.MinLength("password", 8)
		singnUpForm.PasswordMatches("password", "confirmpassword")
		singnUpForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !singnUpForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}
		pExists := service.PhoneExists(r.FormValue("phone"))
		//pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		if *pExists {
			singnUpForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}
		eExists := service.EmailExists(r.FormValue("email"))
		//eExists := uh.userService.EmailExists(r.FormValue("email"))
		if *eExists {
			singnUpForm.VErrors.Add("email", "Email Already Exists")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			singnUpForm.VErrors.Add("password", "Password Could not be stored")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		role,errs := service.RoleByName("USER")
		//role, errs := uh.userRole.RoleByName("USER")

		if err != nil {
			singnUpForm.VErrors.Add("role", "could not assign role to the user")
			uh.tmpl.ExecuteTemplate(w, "signup.layout", singnUpForm)
			return
		}

		user := &entity.User{
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   role.ID,
		}
		_,errs = service.RegisterUser(user)
		//_, errs = uh.userService.StoreUser(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
// Logout hanldes the POST /logout requests
func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(userSess.UUID, w)
	service.DeleteSession(userSess.UUID)
	//uh.sessionService.DeleteSession(userSess.UUID)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
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
func (uh *UserHandler) AdminUsers(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
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
func (uh *UserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
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
			Roles   []entity.Role
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

		pExists := uh.userService.PhoneExists(r.FormValue("phone"))
		if pExists {
			accountForm.VErrors.Add("phone", "Phone Already Exists")
			uh.tmpl.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
			return
		}
		eExists := uh.userService.EmailExists(r.FormValue("email"))
		if eExists {
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
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   uint(roleID),
		}
		_, errs := uh.userService.StoreUser(user)
		if len(errs) > 0 {
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
func (uh *UserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		//idRaw := r.URL.Query().Get("id")
		idRaw := ps.ByName("id")
		id, err := strconv.Atoi(idRaw) 
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		//user, errs := uh.userService.User(uint(id))
		user,errs := service.GetUser(uint(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		//roles, errs := uh.userRole.Roles()
		roles,errs := service.Roles()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		//role, errs := uh.userRole.Role(user.RoleID)
		role,errs := service.Role(user.RoleID)
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
		user,errs := service.GetUser(uint(uid))
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
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: user.Password,
			RoleID:   uint(roleID),
		}
		//_, errs = uh.userService.UpdateUser(usr)
		_,errs = service.EditUser(usr)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersDelete handles Delete /admin/users/delete?id={id} request
func (uh *UserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	if r.Method == http.MethodGet {
		//idRaw := r.URL.Query().Get("id")
		idRaw := ps.ByName("id")
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


