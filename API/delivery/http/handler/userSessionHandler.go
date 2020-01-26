package handler

import (
	"fmt"
	"github.com/goEventListing/API/entity"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/goEventListing/API/user"
)

//UserSessionHandler ...
type UserSessionHandler struct{
	userSessionSrv user.SessionService
}


//NewUserSessionHandler initializes and returns new UserSessionHandler
func NewUserSessionHandler(USS user.SessionService) *UserSessionHandler {
	return &UserSessionHandler{userSessionSrv: USS}
}
//Session ...
func(ush *UserSessionHandler) Session (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	sessionID := ps.ByName("uuid")

	session,errs := ush.userSessionSrv.Session(sessionID)

	if len(errs)>0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(session,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
	return
}
//StoreSession ...
func(ush *UserSessionHandler) StoreSession (w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	l :=req.ContentLength
	body := make([]byte,l)
	req.Body.Read(body)
	session := &entity.Session{}

	err:= json.Unmarshal(body,session)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	sess,errs := ush.userSessionSrv.StoreSession(session)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	p :=fmt.Sprintf("/session/create/%d",session.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	output,err:= json.MarshalIndent(sess,"","\t\t")

	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(output)

	return

}
//DeleteSession ...
func(ush *UserSessionHandler) DeleteSession(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	sessionID := ps.ByName("uuid")

	_,errs := ush.userSessionSrv.DeleteSession(sessionID)
	
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
