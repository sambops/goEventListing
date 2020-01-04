package events

import "github.com/birukbelay/Aprojects/EventListing/entity"

// ReviewRepository repository
type ReviewRepository interface {
	Reviews() ([]entity.Review, []error)

	
	
	MakeReview(review *entity.Review, comment entity.Comment) error
	
	getMyReviews(user *entity.User) ([]entity.Review, []error)

	
	setRating(eventID int) error //this is done in the back

	notify(eventID int, tagsID []int) error

	getMyRating(UID, EventID int) int
	deleteComment(id int) error
	getcomments(eventID int) []entity.Comment

}
