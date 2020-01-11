package repository

// import (
// 	"database/sql"
// 	"errors"

// 	"github.com/goEventListing/API/entity"
// )

// // ReviewRepoImpl implements the review.ReviewRepository interface
// type ReviewRepoImpl struct {
// 	conn *sql.DB
// }

// // NewReviewRepoImpl will create an object of PsqlReviewRepository
// func NewReviewRepoImpl(Con *sql.DB) *ReviewRepoImpl {
// 	return &ReviewRepoImpl{conn: Con}
// }

// // Reviews returns all Reviews from the database
// func (rri *ReviewRepoImpl) Reviews() ([]entity.Review, error) {

// 	query := "SELECT * FROM review"
// 	rows, err := rri.conn.Query(query)
// 	if err != nil {
// 		return nil, errors.New("Could not query the database")
// 	}
// 	defer rows.Close()

// 	rvws := []entity.Review{}

// 	for rows.Next() {
// 		review := entity.Review{}
// 		err = rows.Scan(&review.ID, &review.Rating, &review.UserID, &review.EventID, &review.Message, &review.ReviewedAt)

// 		if err != nil {
// 			return nil, err
// 		}
// 		rvws = append(rvws, review)
// 	}

// 	return rvws, nil
// }

// // GetMyReviews returns the reviews of a single Event
// func (rri *ReviewRepoImpl) EventReviews(id int) ([]entity.Review, error) {

// 	query := "SELECT * FROM review WHERE Event_id = $1"
// 	rows, err := rri.conn.Query(query, id)

// 	if err != nil {
// 		return nil, errors.New("Could not query the database")
// 	}
// 	defer rows.Close()

// 	rvws := []entity.Review{}

// 	for rows.Next() {
// 		review := entity.Review{}
// 		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
// 		if err != nil {
// 			return nil, err
// 		}

// 		rvws = append(rvws, review)
// 	}

// 	return rvws, nil
// }

// // Review returns a Review with a given id
// func (rri *ReviewRepoImpl) Review(id int) (entity.Review, error) {

// 	query := "SELECT * FROM review WHERE id = $1"
// 	row := rri.conn.QueryRow(query, id)

// 	r := entity.Review{}

// 	err := row.Scan(&r.ID, &r.Rating, &r.EventID, &r.UserID, &r.Message, &r.ReviewedAt)
// 	if err != nil {
// 		return r, err
// 	}

// 	return r, nil
// }

// // MakeReview stores new review information to database
// func (rri *ReviewRepoImpl) MakeReview(r entity.Review) error {

// 	query := "INSERT INTO review (rating,event_id,user_id,message) values($1, $2, $3, $4)"
// 	_, err := rri.conn.Exec(query, r.Rating, r.EventID, r.UserID, r.Message)
// 	if err != nil {
// 		return errors.New("reviewing has failed")
// 	}

// 	return nil
// }

// // UpdateReview updates a given object with a new data
// func (rri *ReviewRepoImpl) UpdateReview(r entity.Review) error {
// 	query := "UPDATE review SET rating=$1,event_id=$2, user_id=$3, message=$4 WHERE id=$5"
// 	_, err := rri.conn.Exec(query, r.Rating, r.EventID, r.UserID, r.Message, r.ID)

// 	if err != nil {
// 		return errors.New("updating has failed")
// 	}

// 	return nil
// }

// // DeleteReview removes a review from a database by its id
// func (rri *ReviewRepoImpl) DeleteReview(id int) error {
// 	query := "DELETE FROM review WHERE id=$1"
// 	_, err := rri.conn.Exec(query, id)

// 	if err != nil {
// 		return errors.New("Delete has failed")
// 	}

// 	return nil
// }

// // GetMyReviews returns the reviews of a single user
// func (rri *ReviewRepoImpl) GetMyReviews(id int) ([]entity.Review, error) {

// 	query := "SELECT * FROM review WHERE User_id = $1"
// 	rows, err := rri.conn.Query(query, id)

// 	if err != nil {
// 		return nil, errors.New("Could not query the database")
// 	}
// 	defer rows.Close()

// 	rvws := []entity.Review{}

// 	for rows.Next() {
// 		review := entity.Review{}
// 		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
// 		if err != nil {
// 			return nil, err
// 		}

// 		rvws = append(rvws, review)
// 	}

// 	return rvws, nil
// }

// //SetRating sets the average rating of an event after every reviews
// func (rri *ReviewRepoImpl) SetRating(Eid int) error {

// 	query := "SELECT AVG(rating) FROM review WHERE Event_id = $1"
// 	row := rri.conn.QueryRow(query, Eid)

// 	var rating float32
// 	err := row.Scan(&rating)

// 	if err != nil {
// 		return errors.New("Could not make average in the database")
// 	}
// 	_, er := rri.conn.Exec("UPDATE events SET rating=$1 WHERE id=$1", rating, Eid)

// 	if er != nil {
// 		return errors.New("setting new rating has failed")
// 	}
// 	return nil
// }
