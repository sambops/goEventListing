package handler

import (
	"github.com/birukbelay/Aprojects/goEventListing/entity"
	"github.com/birukbelay/Aprojects/goEventListing/review"
	"html/template"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	tmpl        *template.Template
	ReviewServ Review.ReviewService
}

func NewReviewHandler(T *template.Template, rs Review.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		tmpl:        T,
		ReviewServ: rs
	}
}

// Reviews handle requests on route 
func (rh *ReviewHandler) Reviews(w http.ResponseWriter, r *http.Request) {

	reviews, err := rh.ReviewServ.Reviews()
	if err != nil {
		panic(err)
	}
	rh.tmpl.ExecuteTemplate(w, "", reviews)
}

// Reviews handle requests on route 
func (rh *ReviewHandler) EventReviews(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		reviews, err := rh.ReviewServ.EventReviews(id)

		if err != nil {
			panic(err)
		}

		rh.tmpl.ExecuteTemplate(w, "", reviews)

	}

	http.Redirect(w, r, "..", http.StatusSeeOther)
}

func (rh *ReviewHandler) Review(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		review, err := rh.ReviewServ.Review(id)

		if err != nil {
			panic(err)
		}

		rh.tmpl.ExecuteTemplate(w, "", review)

	}

	http.Redirect(w, r, "..", http.StatusSeeOther)
}

func (rh *ReviewHandler) GetMyReviews(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		reviews, err := rh.ReviewServ.GetMyReviews(id)

		if err != nil {
			panic(err)
		}

		rh.tmpl.ExecuteTemplate(w, "", reviews)

	}

	http.Redirect(w, r, "..", http.StatusSeeOther)
}
// ReviewsNew hanlde requests on route
func (rh *ReviewHandler) NewReview(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		rvw := entity.Review{}
		rvw.Rating = r.FormValue("rating")
		rvw.Message = r.FormValue("message")
		rvw.EventID = r.FormValue("Eid")
		rvw.UserID = r.FormValue("Uid")

		
		err = rh.ReviewServ.MakeReview(rvw)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "...", http.StatusSeeOther)

	} else {

		rh.tmpl.ExecuteTemplate(w, "...", nil)

	}
}


// ReviewsNew hanlde requests on route
func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {

	
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		review, err := rh.ReviewServ.Review(id)

		if err != nil {
			panic(err)
		}

		rh.tmpl.ExecuteTemplate(w, "", review)

	} else if r.Method == http.MethodPost {

		
		rvw := entity.Review{}
		rvw.Rating = r.FormValue("rating")
		rvw.Message = r.FormValue("message")
		rvw.EventID = r.FormValue("Eid")
		rvw.UserID = r.FormValue("Uid")

		

		
		err = rh.ReviewServ.UpdateReview(rvw)

		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}

}


// ReviewsNew hanlde requests on route
func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		
		
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		
		err = rh.ReviewServ.DeleteReview(id)

		if err != nil {
			panic(err)
		}

		

	} 

		http.Redirect(w, r, "...", http.StatusSeeOther)

	
}
	
	
