package review

import "github.com/goEventListing/API/entity"

// ReviewRepository specifies application review related operations
type ReviewRepository interface {
	Reviews() ([]entity.Review, error)
	EventReviews(id int) ([]entity.Review, error)
	Review(id int) (entity.Review, error)
	MakeReview(r entity.Review) error
	SetRating(Eid int) error
	UpdateReview(r entity.Review) error
	DeleteReview(id int) error
	GetMyReviews(id int) ([]entity.Review, error)

	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
