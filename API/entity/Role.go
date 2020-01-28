package entity


// Role repesents application user roles
type Role struct {
	ID    uint `json:"id"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Users []User
}