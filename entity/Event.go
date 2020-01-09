package entity

import (
	"time"

	"github.com/birukbelay/Aprojects/goEventListing/entity"
)

//Event shows event intity
type Event struct {
	ID int

	Name, Details, Image              string
	UserID                            int
	City, Country, Place, Coordinates string

	IsPassed   bool
	Rating     int
	PostedDate time.Time
	price      float32
	Reviews    []entity.Review
	Tags       []entity.Tag
}
