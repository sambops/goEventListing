package repository

import (
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/eventListing/entity"
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *sql.DB
	
}

// NewEventRepoImp will create an object of EventRepoImpl

func NewEventRepoImp(con *sql.DB) *EventRepoImpl {
    return &EventRepoImpl{conn: con}
}
