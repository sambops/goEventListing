package repository

import (
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
)


type ReviewRepoImpl struct {
	conn *sql.DB
	
}



func NewReviewRepoImp(con *sql.DB) *EventRepoImpl {
    return &ReviewRepoImpl{conn: con}
}

