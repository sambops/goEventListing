package handler

import (
	"strconv"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/user"
	
)

//UserHandler handles user related requests
type UserHandler struct {
	userSrv user.UserService
}
//GETUSER


//NewUserHandler initializes and returns new UserHandler
func NewUserHandler(US user.UserService) *UserHandler {
	return &UserHandler{userSrv: US}
}


//GetUser ... handles GET /el/user/:id request
func(uh *UserHandler) GetUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(err)


	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}


	user,err := uh.userSrv.GetUser(uint(id))
	fmt.Println(err)

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println(user)

	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{


		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(err)

	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return

}

//GetUsers ... handles GET /el/users/
func(uh *UserHandler) GetUsers(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	
	usrs,errs := uh.userSrv.GetUsers()
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(usrs,"","\t\t")
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return

}

//GetUserByUserName ... 
func(uh *UserHandler) GetUserByUserName(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	name:= ps.ByName("userName")


	user,err := uh.userSrv.GetUserByUserName(name)
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
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
	
	if err != nil {
		fmt.Println("errorroing 3")
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user,err = uh.userSrv.RegisterUser(user)
	

	if err != nil {
	
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{
		
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	
	// p := fmt.Sprintf("/el/user/register/%d", user.ID)
	// w.Header().Set("Location",p)
	// w.WriteHeader(http.StatusCreated)
	return
}

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

}
//DeleteUser ... handle POST /el/user/remove:id
func(uh *UserHandler) DeleteUser(w http.ResponseWriter,req *http.Request,ps httprouter.Params	){
	id, err := strconv.Atoi(ps.ByName("id"))


	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_,err = uh.userSrv.DeleteUser(uint(id))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
//UserRoles ... handle POST  /el/user/role/:user
func(uh *UserHandler) UserRoles(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	l := req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)

	user := &entity.User{}
	err := json.Unmarshal(body,user)

	fmt.Println(user)
	
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	rols,errs := uh.userSrv.UserRoles(user)
	

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output,err:= json.MarshalIndent(rols,"","\t\t")
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	
	// p := fmt.Sprintf("/el/user/register/%d", user.ID)
	// w.Header().Set("Location",p)
	// w.WriteHeader(http.StatusCreated)
	return
}

//PhoneExists ... handles GET /el/user/check/:phone
func(uh *UserHandler) PhoneExists(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	phone:= ps.ByName("phone")
	exists := uh.userSrv.PhoneExists(phone)

	output,err:= json.MarshalIndent(exists,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return

	
}

//EmailExists ... handles GET /el/user/check/:email
func(uh *UserHandler) EmailExists(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	email:= ps.ByName("email")
	exists := uh.userSrv.EmailExists(email)

	output,err:= json.MarshalIndent(exists,"","\t\t")
	
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return

}

//UserByEmail ... handles GET /el/user/:email
func(uh *UserHandler) UserByEmail(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	email:= ps.ByName("email")

	user,errs := uh.userSrv.UserByEmail(email)
	if len(errs) > 0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output,err:= json.MarshalIndent(user,"","\t\t")
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}





