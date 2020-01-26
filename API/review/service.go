package review

import "github.com/goEventListing/API/entity"

// ReviewService specifies application review related services
type ReviewService interface {
	Reviews() ([]entity.Review, []error)

	Review(id uint) (*entity.Review, []error)
	MakeReviewAndRating(r *entity.Review) (*entity.Review, []error)

	UpdateReview(r *entity.Review) (*entity.Review, []error)
	DeleteReview(id uint) (*entity.Review, []error)
	EventReviews(id uint) ([]entity.Review, []error)
	//GetMyReviews(id uint) ([]entity.Review, []error)
	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
