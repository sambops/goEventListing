package review

import "github.com/birukbelay/Aprojects/goEventListing/entity"

// ReviewRepository specifies application review related operations
type ReviewRepository interface {
	Reviews() ([]entity.Review, error)
	Review(id int) (entity.Review, error)
	MakeReview(r entity.Review) error
	SetRating(Eid int) error
	UpdateReview(r entity.Review) error
	DeleteReview(id int) error
	GetMyReviews(id int) ([]entity.Review, error)
	setRating(eventID int) error //this is done in the back when a event is rated

	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
