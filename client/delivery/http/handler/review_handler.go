package handler

import (
	"strconv"
	"goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
)

//ReviewHandler ... handles review related request
type ReviewHandler struct{
	tmpl *template.Template
}
//NewReviewHandler ... initializes and returns new EventHandler
func NewReviewHandler(T *template.Template) *ReviewHandler{
	return &ReviewHandler{tmpl:T}
}
//MakeReviewAndRating ... handlers request on /el/review/make
func (rh *ReviewHandler) MakeReviewAndRating(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	usr := GetUser(w,req)
	usr, errr := service.GetUser(usr.ID)
	if errr != nil {
			// http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	
		}

		var review *entity.Review
		if req.Method == http.MethodPost{
			//review form value

		}
		rh.tmpl.ExecuteTemplate(w,"template.html",review)
}

//EventReviews ... handles request on /el/review/event/:id
func (rh *ReviewHandler) EventReviews(w http.ResponseWriter,req *http.Request,ps httprouter.Params){
	idraw := req.FormValue("id")

	id, err := strconv.Atoi(idraw)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	usr := GetUser(w,req)
	usr,errr := service.GetUser(usr.ID)

	if errr != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}

	review,err := service.EventReviews(uint(id))
	
	if err != nil{
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	
	rh.tmpl.ExecuteTemplate(w,"template.html",review)
}
//UpdateReview ... handles request on /el/review/edit
func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

	usr := GetUser(w,req)
	usr, errr := service.GetUser(usr.ID)

	if errr != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	var evt *entity.Event
	if req.Method == http.MethodPost{
		//update code
	}
	rh.tmpl.ExecuteTemplate(w,"tempalte.html",evt)


}


