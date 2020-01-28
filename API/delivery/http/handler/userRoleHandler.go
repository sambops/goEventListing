package handler

import (
	"fmt"
	"github.com/goEventListing/API/entity"
	"encoding/json"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/goEventListing/API/user"
)

//UserRoleHandler handles user role realated requests
type UserRoleHandler struct{
	userRoleSrv user.RoleService
	
}


//NewUserRoleHandler initializes and returns new UserHandler
func NewUserRoleHandler(URS user.RoleService) *UserRoleHandler {
	return &UserRoleHandler{userRoleSrv: URS}
}
//Roles ... 
func(urh *UserRoleHandler) Roles (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	role,errs := urh.userRoleSrv.Roles()

	if len(errs)>0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(role,"","\t\t")

	if err!=nil{
		w.Header().Set("Content-Type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
//Role ...
func(urh *UserRoleHandler) Role (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))
	
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	role,errs := urh.userRoleSrv.Role(uint(id))

	if len(errs)>0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(role,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
//RoleByName ...
func(urh *UserRoleHandler) RoleByName (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	name:= ps.ByName("name")

	role,errs := urh.userRoleSrv.RoleByName(name)

	if len(errs)>0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(role,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
//UpdateRole ...
func(urh *UserRoleHandler) UpdateRole (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id,err := strconv.Atoi(ps.ByName("id"))
		
if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
role,errs := urh.userRoleSrv.Role(uint(id))

if len(errs) > 0{
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
l:= req.ContentLength
body := make([]byte,l)
req.Body.Read(body)

json.Unmarshal(body,&role)

role,errs = urh.userRoleSrv.UpdateRole(role)

if len(errs) > 0 {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
output, err := json.MarshalIndent(role, "", "\t\t")

if err != nil {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
w.Header().Set("Content-Type", "application/json")
w.Write(output)
return

}
//DeleteRole ...
func(urh *UserRoleHandler) DeleteRole (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_,errs := urh.userRoleSrv.DeleteRole(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
//StoreRole ...
func(urh *UserRoleHandler) StoreRole(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	l :=req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)
	role := &entity.Role{}

	err:= json.Unmarshal(body,role)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	role,errs := urh.userRoleSrv.StoreRole(role)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	p :=fmt.Sprintf("/role/create/%d",role.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	output,err:= json.MarshalIndent(role,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)

	return
}




