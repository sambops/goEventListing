package repository

import (


	"github.com/goEventListing/API/entity"

	"github.com/jinzhu/gorm"
)

// ReviewGormRepo ...
type ReviewGormRepo struct {
	conn *gorm.DB
}

// NewReviewGormRepo will create an object of PsqlReviewRepository
func NewReviewGormRepo(db *gorm.DB) *ReviewGormRepo {
	return &ReviewGormRepo{conn: db}
}

// Reviews returns all Reviews from the database
func (rgr *ReviewGormRepo) Reviews() ([]entity.Review, []error) {
	review := []entity.Review{}
	errs := rgr.conn.Find(&review).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return review, errs
}

// Review returns a Review with a given id
func (rgr *ReviewGormRepo) Review(id uint) (*entity.Review, []error) {
	rvw := entity.Review{}
	errs := rgr.conn.First(&rvw, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &rvw, errs
}

// UpdateReview updates a given object with a new data
func (rgr *ReviewGormRepo) UpdateReview(r *entity.Review) (*entity.Review, []error) {
	rvw := r
	errs := rgr.conn.Save(rvw).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rvw, errs
}

// DeleteReview removes a review from a database by its id
func (rgr *ReviewGormRepo) DeleteReview(id uint) (*entity.Review, []error) {

	rvw, errs := rgr.Review(id)
	if len(errs) > 0 {
		return nil, errs
	}

	errs = rgr.conn.Delete(rvw, rvw.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return rvw, errs
}

//MakeReviewAndRating stores new review information to database
func (rgr *ReviewGormRepo) MakeReviewAndRating(rev *entity.Review) (*entity.Review, []error) {
	rvw := rev
	errs := rgr.conn.Create(rvw).GetErrors()
	if len(errs) > 0 {
		//fmt.Println("MakeReview-err len>1", errs)
		return nil, errs
	}
	return rvw, errs
}

// // GetMyReviews returns the reviews of a single user
// func (rgr *ReviewGormRepo) GetMyReviews(id uint) ([]entity.Review, []error) {
// 	reviews := []entity.Review{}
// 	errs := rgr.conn.Where("user_id = ?", id).Find(&reviews).GetErrors()
// 	if len(errs) > 0 {
// 		//fmt.Println("", errs)
// 		return nil, errs
// 	}
// 	// fmt.Println(review)
// 	return reviews, errs
// }

// EventReviews returns the reviews of a single Event
func (rgr *ReviewGormRepo) EventReviews(id uint) ([]entity.Review, []error) {

	review := []entity.Review{}
	errs := rgr.conn.Where("event_id = ?", id).Find(&review).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	// fmt.Println(review)
	return review, errs

}
