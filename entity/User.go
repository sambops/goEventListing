package entity

//User ... represents users of our system
type User struct {
	//UserID    int
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  []byte
	Phone     string
	Image     string
}
// CREATE TABLE users(
// 	user_id serial PRIMARY KEY,
// 	username VARCHAR (50) UNIQUE NOT NULL,
// 	first_name VARCHAR (50) NOT NULL,
// 	last_name VARCHAR (50) NOT NULL,
// 	email VARCHAR (50) UNIQUE NOT NULL,
// 	password TEXT NOT NULL,
// 	phone VARCHAR (50),
// 	image TEXT
//  );