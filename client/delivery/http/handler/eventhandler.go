package handler

import (
	"fmt"
	"strconv"

	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/service"
	"github.com/julienschmidt/httprouter"

	"html/template"

	"net/http"
)

//UserHandler handles user related requests
type EventHandler struct {
	tmpl *template.Template
}

//NewUserHandler initializes and returns new UserHandler
func NewEventHandler(T *template.Template) *EventHandler {
	return &EventHandler{tmpl: T}
}

func (eh *EventHandler) Event(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	type Data struct {
		event   entity.Event
		reviews []entity.Review
	}

	if r.Method == http.MethodGet {
		fmt.Println("---handler--metd is get--") //metd is get
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		evnt := &entity.Event{}
		revws := []entity.Review{}

		evnt, err = service.Event(uint(id))

		fmt.Println("---handler----got evnt---", evnt) //handler----got evnt
		revws, err = service.EventReviews(uint(id))
		fmt.Println("---handler----got revws---", revws)
		fmt.Println(evnt, evnt)
		evnt.Reviews = revws
		fmt.Println("---handler-- evnts field set--event.revws---", evnt)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// u = &entity.Event{FirstName: fn, LastName: ln, UserName: un, Email: email, Password: bs, Phone: phone, Image: img}

		eh.tmpl.ExecuteTemplate(w, "event.single", evnt)
	}
	//redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return

}
