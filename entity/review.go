package entity

import (
	"time"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
)

// Review is when a user rates to an event
type Review struct {
	ID      int
	Rating  int
	Event   *entity.Event
	User    *entity.User
	Message string

	ReviewedAt time.Time
	// isempty    bool
}
