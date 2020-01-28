package entity

//Tag there are multiple tags for the user to choose from
type Tag struct {
	ID int `json:"id"`
	Name string `json:"name" gorm:"type:varchar(255);not null;unique"` 
	//description string `gorm:"type:text;not null"` 
	//icon        string `gorm:"type:varchar(255);not null"` 
	//User []User `gorm:"many2many:user_tag"`
	//Event []Event `gorm:"many2many:event_tag"`

}

// //UserTags is used when a user register to tags
// type UserTags struct {
// 	UserID, TagID int
// }

// //EventTags is used when an event is posted
// type EventTags struct {
// 	Event_ID, TagID int
// }
