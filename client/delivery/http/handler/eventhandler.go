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

//EventHandler handles user related requests
type EventHandler struct {
	tmpl *template.Template
}

//NewEventHandler initializes and returns new UserHandler
func NewEventHandler(T *template.Template) *EventHandler {
	return &EventHandler{tmpl: T}
}

// Event ...
func (eh *EventHandler) Event(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	type Data struct {
		event   entity.Event
		reviews []entity.Review
	}

	if r.Method == http.MethodGet {

		id, err := strconv.Atoi(ps.ByName("id"))

		fmt.Println("@@---handler/eventhandler/line:38 \n--metd is get--id ==", id) //metd is get
		evnt := &entity.Event{}
		// revws := []entity.Review{}

		evnt, err = service.Event(uint(id))

		fmt.Println("@@--------handler/eventhandler/line:44 \n-- event==", evnt, "\n ###--handler-ended") //handler----got evnt
		// revws, err = service.EventReviews(uint(id))
		// fmt.Println("---handler----got revws---", revws)
		// fmt.Println(evnt, evnt)
		// evnt.Reviews = revws
		// fmt.Println("---handler-- evnts field set--event.revws---", evnt)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("@@-rederict evnt fr")
			return
		}

		// u = &entity.Event{FirstName: fn, LastName: ln, UserName: un, Email: email, Password: bs, Phone: phone, Image: img}
		fmt.Println("template")
		eh.tmpl.ExecuteTemplate(w, "event.single", evnt)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//redirect

}
