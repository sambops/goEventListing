package entity

import "time"

// Notification whenbevent is posted it takes TagID from event_tag & match it with user_tags and inserts into the notification table
type Notification struct {
	ID              int
	EventID, UserID int
	status          bool      //to show that it is seenor not there must be a tigger when the user opens it it will turn it to false
	EndDate         time.Time //
}
