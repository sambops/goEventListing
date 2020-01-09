package entity

//Tag there are multiple tags for the user to choose from
type Tag struct {
	ID          int
	Name        string
	Description string
	Icon        string
}

//User_tags is used when a user register to tags
// type User_tags struct {
// 	UserID, TagID int
// }

// Event_Tag is used when an event is posted
// type Event_Tag struct {
// 	Event_ID, TagID int
// }
