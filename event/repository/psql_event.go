package repository

import (
	"github.com/birukbelay/Aprojects/goEventListing/event"
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *sql.DB
	
}

// NewEventRepoImp will create an object of EventRepoImpl

func NewEventRepoImp(con *sql.DB) *EventRepoImpl {
    return &EventRepoImpl{conn: con}
}

//Event ..returns single event
func (eri *EventRepoImpl) Event(id int) (entity.Event, error) {

	event := entity.Event{}
	event.Reviews =[]entity.Review{}
	event.Tags=[]Tag{}

	query:="SELECT * FROM events WHERE id = $1"
    row := eri.conn.QueryRow(query, id)  
    err := rows.Scan(&event.EId, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.host, &event.IsPassed, &event.Rating, &event.PostedDate, &event.price)
    if err != nil {
        return event, err
    }
	
	query2 := "SELECT * FROM review WHERE Event_id = $1"
	rows, err := rri.conn.Query(query2, id)

	if err != nil {
		return nil, errors.New("Could not query the reviews database")
	}
	defer rows.Close()	
	// rvws := []entity.Review{}
	for rows.Next() {
		review := entity.Review{Event:&event} //it is refering to the event for its *event field
		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.UserID, &review.EventID, &review.Message)
		if err != nil {
			return nil, err
		}

		event.Reviews = append(event.Reviews, review)
	}

	
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



    return event, nil
}





















