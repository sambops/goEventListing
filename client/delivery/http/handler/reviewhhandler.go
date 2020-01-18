package handler

import (
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"github.com/julienschmidt/httprouter"

	"html/template"

	"strconv"

	"net/http"
	// uuid "github.com/satori/go.uuid"
)

//UserHandler handles user related requests
type Reviewhandler struct {
	tmpl *template.Template
}

//NewUserHandler initializes and returns new UserHandler
func NewReviewhandler(T *template.Template) *Reviewhandler {
	return &Reviewhandler{tmpl: T}
}
func (rh *Reviewhandler) MakeReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodPost {

		rvw := &entity.Review{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		rvw.ID = uint(id)
		rvw.Message = r.FormValue("message")

		rate, _ := strconv.Atoi(r.FormValue("rating"))
		rvw.Rating = int(rate)
		eid, _ := strconv.Atoi(r.FormValue("Eid"))
		rvw.EventID = uint(eid)
		uid, _ := strconv.Atoi(r.FormValue("Uid"))
		rvw.UserID = uint(uid)

		review, err := service.MakeReview(rvw)

		if err != nil {
			http.Redirect(w, r, "el/event", http.StatusSeeOther)
			return
		}

		// u = &entity.Event{FirstName: fn, LastName: ln, UserName: un, Email: email, Password: bs, Phone: phone, Image: img}
		rh.tmpl.ExecuteTemplate(w, "signup.html", review)
	}
	//redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return

}
