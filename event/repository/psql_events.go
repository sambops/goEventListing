package repository

import (
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
	// _ "github.com/lib/pq"
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *sql.DB
}

// NewEventRepoImp will create an object of EventRepoImpl
func NewEventRepoImp(Con *sql.DB) *EventRepoImpl {
	return &EventRepoImpl{conn: Con}
}

// Events returns all Events from the database
func (eri *EventRepoImpl) Events() ([]entity.Event, error) {

	rows, err := eri.conn.Query("SELECT * FROM events;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	events := []entity.Event{}

	for rows.Next() {
		event := entity.Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.UserID, &event.IsPassed, &event.Rating, &event.PostedDate, &event.Price)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

//Event ..returns single event
func (eri *EventRepoImpl) Event(id int) (entity.Event, error){

	event :=entity.Event{}
	
	event.Reviews =[]entity.Review{}
	event.Tags=[]entity.Tag{}

	query:="SELECT * FROM events WHERE id = $1"
	row:= eri.conn.QueryRow(query, id)

	e:= row.Scan(&event.ID, &event.Name, &event.Details, &event.Image, &event.Country, &event.City, &event.Date, &event.Place, &event.Coordinates, &event.UserID, &event.IsPassed, &event.Rating, &event.PostedDate, &event.Price)

	if e!= nil{
		return event, e
	}

	// getting the reviews for that event
	query2 := "SELECT * FROM review WHERE Event_id = $1"
	rows, er := eri.conn.Query(query2, id)

	if er != nil {
		return event, errors.New("Could not query the reviews database")
	}
	defer rows.Close()
	// rvws := []entity.Review{}
	for rows.Next() {
		review := entity.Review{} //it is refering to the event for its *event field
		errs = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
		if errs != nil {
			return nil, errs
		}

		event.Reviews = append(event.Reviews, review)
	}

	// getting tags for that event
	query3 := "SELECT * FROM event_tag WHERE event_id = $1"
	rows, err := rri.conn.Query(query3, id)

	if err != nil {
		return nil, errors.New("Could not query the tags database")
	}
	defer rows.Close()
	// rvws := []entity.Review{}
	for rows.Next() {
		tag := entity.Tag{}
		err = rows.Scan(&Tag.ID, &Tag.Name, &Tag.Description, &Tag.Icon)
		if err != nil {
			return nil, err
		}

		event.Tags = append(event.Tags, tag)
	}

	return event, nil

}
