package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
	"github.com/birukbelay/Aprojects/goEventListing/review"
)

// ReviewHandler ...
type ReviewHandler struct {
	tmpl       *template.Template
	ReviewServ review.ReviewService
}

// NewReviewHandler ...
func NewReviewHandler(T *template.Template, rs review.ReviewService) *ReviewHandler {
	return &ReviewHandler{tmpl: T, ReviewServ: rs}
}

// Reviews handle requests on route
func (rh *ReviewHandler) Reviews(w http.ResponseWriter, r *http.Request) {

	reviews, err := rh.ReviewServ.Reviews()
	if err != nil {
		panic(err)
	}
	rh.tmpl.ExecuteTemplate(w, "", reviews)
}

// EventReviews handle requests on route
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

// Review ...
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

// GetMyReviews ...
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

// NewReview hanlde requests on route
func (rh *ReviewHandler) NewReview(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		rvw := entity.Review{}
		var ratingRaw = r.FormValue("rating")

		var rating, err = strconv.Atoi(ratingRaw)
		if err != nil {
			panic(err)
		}
		rvw.Rating = rating

		rvw.Message = r.FormValue("message")

		var EidRaw = r.FormValue("Eid")

		var Eid, er = strconv.Atoi(EidRaw)
		if er != nil {
			panic(err)
		}
		rvw.EventID = Eid

		var UIDRaw = r.FormValue("Uid")
		var UID, e = strconv.Atoi(UIDRaw)
		if e != nil {
			panic(err)
		}
		rvw.UserID = UID

		err = rh.ReviewServ.MakeReview(rvw)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "...", http.StatusSeeOther)

	} else {

		rh.tmpl.ExecuteTemplate(w, "...", nil)

	}
}

//UpdateReview ...
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

		var ratingRaw = r.FormValue("rating")

		var rating, err = strconv.Atoi(ratingRaw)
		if err != nil {
			panic(err)
		}
		rvw.Rating = rating

		rvw.Message = r.FormValue("message")

		var EidRaw = r.FormValue("Eid")

		var Eid, er = strconv.Atoi(EidRaw)
		if er != nil {
			panic(err)
		}
		rvw.EventID = Eid

		var UIDRaw = r.FormValue("Uid")
		var UID, e = strconv.Atoi(UIDRaw)
		if e != nil {
			panic(err)
		}
		rvw.UserID = UID

		err = rh.ReviewServ.UpdateReview(rvw)

		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	}

}

// DeleteReview ..
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
