package reviews

import "github.com/birukbelay/Aprojects/goEventListing/entity"

// ReviewRepository repository
type ReviewRepository interface {
	Reviews() ([]entity.Review, []error)
	
	MakeReview(review *entity.Review, comment entity.Comment) error
	setRating(eventID int) error //this is done in the back when a event is rated

	deleteReview(id int) error
	updateReview(id int) error

	getMyReviews(user *entity.User) ([]entity.Review, []error) // list of reviews a user have given

	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
