package service

import (
	"github.com/birukbelay/Aprojects/goEventListing/entity"
	"github.com/birukbelay/Aprojects/goEventListing/review"
)

// ReviweServiceImpl implements review.Reviewservice interface
type ReviweServiceImpl struct {
	ReviewRepo review.ReviewRepository
}

// NewReviewServiceImpl will create new CategoryService object
func NewReviewServiceImpl(RevRepo review.ReviewRepository) *ReviweServiceImpl {
	return &ReviweServiceImpl{ReviewRepo: RevRepo}
}

// Reviews returns list of Reviews
func (rs *ReviweServiceImpl) Reviews() ([]entity.Review, error) {

	rvs, err := rs.ReviewRepo.Reviews()

	if err != nil {
		return nil, err
	}

	return rvs, nil
}

// GetEventReviews returns list of Reviews for an event
func (rs *ReviweServiceImpl) EventReviews(id int) ([]entity.Review, error) {

	rvs, err := rs.ReviewRepo.GetEventReviews(id)

	if err != nil {
		return nil, err
	}

	return rvs, nil
}

// Review returns a Review with a given id
func (rs *ReviweServiceImpl) Review(id int) (entity.Review, error) {

	rvw, err := rs.ReviewRepo.Review(id)

	if err != nil {
		return rvw, err
	}

	return rvw, nil

}

// MakeReview stores new review information
func (rs *ReviweServiceImpl) MakeReview(r entity.Review) error {

	err := rs.ReviewRepo.MakeReview(r)

	if err != nil {
		return err
	}
	er := rs.ReviewRepo.SetRating(r.EventID)

	if er != nil {
		return er
	}

	return nil
}

// UpdateReview updates a given review with a new data
func (rs *ReviweServiceImpl) UpdateReview(r entity.Review) error {

	err := rs.ReviewRepo.UpdateReview(r)

	if err != nil {
		return err
	}

	return nil
}

// DeleteReview removes a review by its id
func (rs *ReviweServiceImpl) DeleteReview(id int) error {

	err := rs.ReviewRepo.DeleteReview(id)

	if err != nil {
		return err
	}

	return nil
}

// GetMyReviews returns the reviews of a single user
func (rs *ReviweServiceImpl) GetMyReviews(id int) ([]entity.Review, error) {

	reviews, err := rs.ReviewRepo.GetMyReviews(id)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
