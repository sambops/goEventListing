package entity

//Tag there are multiple tags for the user to choose from
type Tag struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description string `json:"description" gorm:"type:text; not null"`
	Icon        string `json:"icon" gorm:"type:varchar(255);not null"`
}

// //UserTags is used when a user register to tags
// type UserTags struct {
// 	UserID, TagID int
// }

// //EventTags is used when an event is posted
// type EventTags struct {
// 	Event_ID, TagID int
// }
