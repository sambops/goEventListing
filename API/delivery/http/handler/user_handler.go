package handler

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/user"
	"github.com/julienschmidt/httprouter"
=======
	"strconv"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/user"
	



>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
)

//UserHandler handles user related requests
type UserHandler struct {
	userSrv user.UserService
}
<<<<<<< HEAD

//GETUSER

=======
//GETUSER


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(US user.UserService) *UserHandler {
	return &UserHandler{userSrv: US}
}

<<<<<<< HEAD
//GetUser ... handles GET /el/user/:id request
func (uh *UserHandler) GetUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//(id unit) (*entity.User, error)
	fmt.Println("here....")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(err)

=======

//GetUser ... handles GET /el/user/:id request
func(uh *UserHandler) GetUser (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	//(id unit) (*entity.User, error)	
	fmt.Println("here....")
	id,err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(err)


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

<<<<<<< HEAD
	user, err := uh.userSrv.GetUser(uint(id))
	fmt.Println(err)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
=======

	user,err := uh.userSrv.GetUser(uint(id))
	fmt.Println(err)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	}

	fmt.Println(user)

<<<<<<< HEAD
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
=======
	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println(err)

<<<<<<< HEAD
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//GetUserByUserName ...
func (uh *UserHandler) GetUserByUserName(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//(id unit) (*entity.User, error)

	name := ps.ByName("userName")

	user, err := uh.userSrv.GetUserByUserName(name)
	if err != nil {
=======
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
	

}
//GetUserByUserName ... 
func(uh *UserHandler) GetUserByUserName (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	//(id unit) (*entity.User, error)

	name:= ps.ByName("userName")


	user,err := uh.userSrv.GetUserByUserName(name)
	if err != nil{
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

<<<<<<< HEAD
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
=======
	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
<<<<<<< HEAD
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//RegisterUser ... handle POST /el/user/register/:user   ....
func (uh *UserHandler) RegisterUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//RegisterUser(user *entity.User)(*entity.User,error)
	l := req.ContentLength
	body := make([]byte, l)
	req.Body.Read(body)

	user := &entity.User{}
	err := json.Unmarshal(body, user)

	fmt.Println(user)

=======
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
	
}
//RegisterUser ... handle POST /el/user/register/:user   ....
func(uh *UserHandler) RegisterUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	//RegisterUser(user *entity.User)(*entity.User,error)
	l := req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)

	user := &entity.User{}
	err := json.Unmarshal(body,user)

	fmt.Println(user)
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if err != nil {
		fmt.Println("errorroing 3")
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
<<<<<<< HEAD
	user, err = uh.userSrv.RegisterUser(user)

	if err != nil {

=======
	user,err = uh.userSrv.RegisterUser(user)
	

	if err != nil {
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

<<<<<<< HEAD
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {

=======
	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{
		
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

<<<<<<< HEAD
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

=======
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	// p := fmt.Sprintf("/el/user/register/%d", user.ID)
	// w.Header().Set("Location",p)
	// w.WriteHeader(http.StatusCreated)
	return
}
<<<<<<< HEAD

//AuthenticateUser ... handle POST /el/user/login/
func (uh *UserHandler) AuthenticateUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//AuthenticateUser(userName string, password string) (*entity.User, error)

	l := req.ContentLength
	body := make([]byte, l)
	req.Body.Read(body)

=======
//AuthenticateUser ... handle POST /el/user/login/
func(uh *UserHandler) AuthenticateUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
//AuthenticateUser(userName string, password string) (*entity.User, error)

	l := req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	authenticate := &entity.Authenticate{}

	//fmt.Println("check here")

<<<<<<< HEAD
	err := json.Unmarshal(body, authenticate)

=======
	err := json.Unmarshal(body,authenticate)
	
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if err != nil {
		//fmt.Println("i'm here")
		//fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	//fmt.Println("check check")

<<<<<<< HEAD
	usr, err := uh.userSrv.AuthenticateUser(authenticate.UserName, authenticate.Password)

	if err != nil {
		//fmt.Print("check me again..")
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(usr, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//EditUser ... handle PUT /el/user/edit:id
func (uh *UserHandler) EditUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//EditUser(user *entity.User)(*entity.User,[]error)

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user, err := uh.userSrv.GetUser(uint(id))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := req.ContentLength
	body := make([]byte, l)
	req.Body.Read(body)

	json.Unmarshal(body, &user)

	user, errs := uh.userSrv.EditUser(user)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//DeleteUser ... handle POST /el/user/remove:id
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

=======
	usr,err:=uh.userSrv.AuthenticateUser(authenticate.UserName,authenticate.Password)
	
	if err != nil{
		//fmt.Print("check me again..")
		fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
	}

	output,err:= json.MarshalIndent(usr,"","\t\t")
	

	if err != nil{
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
	}
=======
//EditUser ... handle POST /el/user/edit:id
func(uh *UserHandler) EditUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
//EditUser(user *entity.User)(*entity.User,[]error)

id,err := strconv.Atoi(ps.ByName("id"))

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
user,err := uh.userSrv.GetUser(uint(id))

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

l :=req.ContentLength
body := make([]byte,l)
req.Body.Read(body)

json.Unmarshal(body, &user)

user,errs := uh.userSrv.EditUser(user)
if len(errs) > 0 {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
output, err := json.MarshalIndent(user, "", "\t\t")

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
w.Header().Set("Content-Type", "application/json")
w.Write(output)
return


//EditUser ... handle PUT /el/user/edit:id
func(uh *UserHandler) EditUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
//EditUser(user *entity.User)(*entity.User,[]error)

id,err := strconv.Atoi(ps.ByName("id"))

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
user,err := uh.userSrv.GetUser(uint(id))

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

l :=req.ContentLength
body := make([]byte,l)
req.Body.Read(body)

json.Unmarshal(body, &user)

user,errs := uh.userSrv.EditUser(user)
if len(errs) > 0 {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
output, err := json.MarshalIndent(user, "", "\t\t")

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
w.Header().Set("Content-Type", "application/json")
w.Write(output)
return

}
//DeleteUser ... handle POST /el/user/remove:id
func(uh *UserHandler) DeleteUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id, err := strconv.Atoi(ps.ByName("id"))


>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
<<<<<<< HEAD
	_, err = uh.userSrv.DeleteUser(uint(id))
=======
	_,err = uh.userSrv.DeleteUser(uint(id))
>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
<<<<<<< HEAD
=======






>>>>>>> 4f0152ae7f3c892c7aff7d17d68061483d53f238
