package entity

import "time"

//Comment shows the coment the user posted on an event
type Comment struct {
	ID int

	UserID, EventID int
	Body            string
	Official        bool
	CommentedAt     time.Time
	
}
