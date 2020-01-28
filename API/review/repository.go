package review

import "github.com/goEventListing/API/entity"

// ReviewRepository specifies application review related operations
type ReviewRepository interface {
	Reviews() ([]entity.Review, []error)

	Review(id uint) (*entity.Review, []error)
	MakeReviewAndRating(r *entity.Review) (*entity.Review, []error)

	UpdateReview(r *entity.Review) (*entity.Review, []error)
	DeleteReview(id uint) (*entity.Review, []error)
	
	//reviews specific to event
	EventReviews(id uint) ([]entity.Review, []error)


	//GetMyReviews(id uint) ([]entity.Review, []error)
	
	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
