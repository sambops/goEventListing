package entity

//Session represents login user session
type Session struct {
	ID         uint `json:"id"`
	UUID       string `json:"uuid" gorm:"type:varchar(255);not null"`
	Expires    int64  `json:"expires" gorm:"type:varchar(255);not null"`
	SigningKey []byte `json:"signingkey" gorm:"type:varchar(255);not null"`
}
