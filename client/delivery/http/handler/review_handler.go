 package handler

import (
	"strconv"
	"github.com/goEventListing/client/entity"
	"github.com/goEventListing/client/form"
	"net/url"
	"github.com/goEventListing/client/rtoken"
	
	"github.com/goEventListing/client/service"
	"net/http"
	"html/template"
)

//ReviewHandler ... handles review related request
type ReviewHandler struct{
	tmpl *template.Template
	csrfSignKey []byte
	ushnd   UserHandler
}
//NewReviewHandler ... initializes and returns new EventHandler
func NewReviewHandler(T *template.Template,csKey []byte) *ReviewHandler{
	return &ReviewHandler{tmpl:T,csrfSignKey: csKey}
}
//MakeReviewAndRating ... handlers request on /el/review/make
func (rh *ReviewHandler) MakeReviewAndRating(w http.ResponseWriter,req *http.Request){
	id := rh.ushnd.loggedInUser.ID

	token, err := rtoken.CSRFToken(rh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if req.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		rh.tmpl.ExecuteTemplate(w, "newEvent.html", newCatForm)

		if req.Method == http.MethodPost {
			// Parse the form data
			err := req.ParseForm()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// Validate the form contents
			newRvwForm := form.Input{Values: req.PostForm, VErrors: form.ValidationErrors{}}
			//newRvwForm.Required("name", "details","country","city","place","price")
			newRvwForm.MinLength("message", 5)
			newRvwForm.CSRF = token
			// If there are any errors, redisplay the signup form.
			if !newRvwForm.Valid() {
				rh.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newRvwForm)
				return
			}
			rating,_  := strconv.Atoi(req.FormValue("rating"))
			eventid,_ := strconv.Atoi(req.URL.Query().Get("id"))
			
		
	
			rvw := &entity.Review{
				Rating: rating,
				UserID : id,
				EventID :uint (eventid),
				Message : req.FormValue("message"),
			}

			//writeFile(&mf, fh.Filename)
			//_, errs := ach.categorySrv.StoreCategory(ctg)
			_,err = service.MakeReviewAndRating(rvw)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			http.Redirect(w, req, "/", http.StatusSeeOther)
			}
		}
}

//EventReviews ... handles request on /el/review/event/:id
func (rh *ReviewHandler) EventReviews(w http.ResponseWriter,req *http.Request){                                                                                        
	// idraw := req.FormValue("id")
	// id, err := strconv.Atoi(idraw)
	// if err != nil {
	// 	w.WriteHeader(http.StatusNoContent)
	// }
	// usr := GetUser(w,req)
	// usr,errr := service.GetUser(usr.ID)
	// if errr != nil {
	// 	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	// }
	// review,err := service.EventReviews(uint(id))
	// if err != nil{
	// 	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	// }
	// rh.tmpl.ExecuteTemplate(w,"template.html",review)
	//event_id
	eventid,_ := strconv.Atoi(req.URL.Query().Get("id"))
	review,err := service.EventReviews (uint (eventid))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	token, err := rtoken.CSRFToken(rh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		reviews []entity.Review
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		reviews : *review,
		CSRF:       token,
	}
	rh.tmpl.ExecuteTemplate(w, "create.html", tmplData)

}
// UpdateReview ... handles request on /el/review/edit
func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter,r *http.Request){
	// 	token, err := rtoken.CSRFToken(rh.csrfSignKey)
	// 	if err != nil {
	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	}
	// 	if r.Method == http.MethodGet {
	// 		idRaw := r.URL.Query().Get("id")
	// 		id, err := strconv.Atoi(idRaw)
	// 		if err != nil {
	// 			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 		}
	// 		rvw,err := service.EventReviews(uint (id))
			
	// 		//cat, errs := rh.categorySrv.Category(uint(id))
	// 		if err != nil {
	// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 		}
			
	// 		values := url.Values{}
	// 		values.Add("catid", idRaw)
	// 		values.Add("catname",rvw. )
	// 		values.Add("catdesc", cat.Description)
	// 		values.Add("catimg", cat.Image)
	// 		upCatForm := struct {
	// 			Values   url.Values
	// 			VErrors  form.ValidationErrors
	// 			Category *entity.Category
	// 			CSRF     string
	// 		}{
	// 			Values:   values,
	// 			VErrors:  form.ValidationErrors{},
	// 			Category: cat,
	// 			CSRF:     token,
	// 		}
	// 		ach.tmpl.ExecuteTemplate(w, "reviewUpdate.html", upCatForm)
	// 		return
	// 	}
	// 	if r.Method == http.MethodPost {
	// 		// Parse the form data
	// 		err := r.ParseForm()
	// 		if err != nil {
	// 			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 			return
	// 		}
	// 		// Validate the form contents
	// 		updateRvwForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
	// 		updateRvwForm.MinLength("message", 5)
	// 		updateRvwForm.CSRF = token

	// 		rvwID, err := strconv.Atoi(r.FormValue("id"))
	// 		if err != nil {
	// 			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 		}
	// 		rvww := &entity.Review{
	// 			ID:          uint (rvwID),
	// 			UserID : id,
	// 			EventID:id,
	// 			Message:r.FormValue("message"),
	// 		}
	// 		// mf, fh, err := r.FormFile("catimg")
	// 		// if err == nil {
	// 		// 	ctg.Image = fh.Filename
	// 		// 	err = writeFile(&mf, ctg.Image)
	// 		// }
	// 		// if mf != nil {
	// 		// 	defer mf.Close()
	// 		// }
	// 		//_, errs := ach.categorySrv.UpdateCategory(ctg)
	// 		_,err = service.UpdateReview(rvww)
	// 		if err != nil {
	// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
	// 		return
	// 	}


}


//DeleteReview ... handles request on /el/review/delete
func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_,err = service.DeleteReview(uint (id))
		//_, errs := ach.categorySrv.DeleteCategory(uint(id))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
