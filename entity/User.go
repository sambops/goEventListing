package entity

//User ... represents users of our system
type User struct {
	UserID    int
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  []byte
	Phone     string
	Image     string
}
