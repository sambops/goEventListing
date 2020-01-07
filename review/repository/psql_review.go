package repository

import (
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
)

// ReviewRepoImpl implements the review.ReviewRepository interface
type ReviewRepoImpl struct {
	conn *sql.DB
}

// NewReviewRepoImpl will create an object of PsqlReviewRepository
func NewReviewRepoImpl(Con *sql.DB) *ReviewRepoImpl {
	return &ReviewRepoImpl{conn: Con}
}

// Reviews returns all Reviews from the database
func (rri *ReviewRepoImpl) Reviews() ([]entity.Review, error) {
	rows, err := rri.conn.Query("SELECT * FROM review;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	rvws := []entity.Review{}

	for rows.Next() {
		review := entity.Review{}
		err = rows.Scan(&review.ID, &review.Rating, &review.UserID, &review.EventID, &review.Message, &review.ReviewedAt)

		if err != nil {
			return nil, err
		}
		rvws = append(rvws, review)
	}

	return rvws, nil
}

// Review returns a Review with a given id
func (rri *ReviewRepoImpl) Review(id int) (entity.Review, error) {

	row := rri.conn.QueryRow("SELECT * FROM review WHERE id = $1", id)

	r := entity.Review{}

	err := row.Scan(&r.ID, &r.Rating, &r.EventID, &r.UserID, &r.Message, &r.ReviewedAt)
	if err != nil {
		return r, err
	}

	return r, nil
}

// MakeReview stores new review information to database
func (rri *ReviewRepoImpl) MakeReview(r entity.Review) error {

	_, err := rri.conn.Exec("INSERT INTO review (rating,event_id,user_id,message) values($1, $2, $3, $4)", r.Rating, r.EventID, r.UserID, r.Message)
	if err != nil {
		return errors.New("reviewing has failed")
	}

	return nil
}

// UpdateReview updates a given object with a new data
func (rri *ReviewRepoImpl) UpdateReview(r entity.Review) error {

	_, err := rri.conn.Exec("UPDATE review SET rating=$1,event_id=$2, user_id=$3, message=$4 WHERE id=$5", r.Rating, r.EventID, r.UserID, r.Message, r.ID)

	if err != nil {
		return errors.New("updating has failed")
	}

	return nil
}

// DeleteReview removes a review from a database by its id
func (rri *ReviewRepoImpl) DeleteReview(id int) error {

	_, err := rri.conn.Exec("DELETE FROM review WHERE id=$1", id)

	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// GetMyReviews returns the reviews of a single user
func (rri *ReviewRepoImpl) GetMyReviews(id int) ([]entity.Review, error) {

	rows, err := rri.conn.Query("SELECT * FROM review WHERE User_id = $1", id)

	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	rvws := []entity.Review{}

	for rows.Next() {
		review := entity.Review{}
		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
		if err != nil {
			return nil, err
		}
		rvws = append(rvws, review)
	}

	return rvws, nil
}

//SetRating sets the average rating of an event after every reviews
func (rri *ReviewRepoImpl) SetRating(Eid int) error {

	row := rri.conn.QueryRow("SELECT AVG(rating) FROM review WHERE Event_id = $1", Eid)

	var rating float32
	err := row.Scan(&rating)

	if err != nil {
		return errors.New("Could not make average in the database")
	}
	_, err := rri.conn.Exec("UPDATE events SET rating=$1 WHERE id=$1", rating, Eid)

	if err != nil {
		return errors.New("setting new rating has failed")
	}
	return nil
}
