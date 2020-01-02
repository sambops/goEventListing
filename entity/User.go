package entity

<<<<<<< HEAD
//User users participate in events
type User struct {
	ID                                          int
	FName, LName, Username, Email, Phone, Image string
=======
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
>>>>>>> e18614362e5300c66820568db16e16c72c4c3f76
}
