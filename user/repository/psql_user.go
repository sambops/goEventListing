package repository

import (
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"errors"

	"github.com/EventListing/entity"
)

//UserRepositoryImpl ... implements the User.UserRepository interface
type UserRepositoryImpl struct {
	conn *sql.DB
}

//NewUserRepositoryImpl will create an object of  UserReposiotryImpl
func NewUserRepositoryImpl(Conn *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}

//Methods with pointer receivers can modify the value to which the receiver points.
//Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

//RegisterUser ... this is a method to register our users int to the user table
func (uri *UserRepositoryImpl) RegisterUser(user entity.User) error {

	//username taken?
	_, err := uri.conn.Query("SELECT * FROM users where UserName = $1", user.UserName)
	if err == nil {
		return errors.New("user name already taken try other")
	}
	_, err = uri.conn.Exec("INSERT INTO users(FirstName,LastName,UserName,Email,Password,Phone,Image) values($1,$2,$3,$4,$5,$6,$7", user.FirstName, user.LastName, user.UserName, user.Email,user.Password, user.Phone, user.Image)

	if err != nil {
		return errors.New("insertion has failed")
	}
	return nil

}

//AuthenticateUser ... this is a  method to authenticate a user before logining in
func (uri *UserRepositoryImpl) AuthenticateUser(userName string, password string) (entity.User, error) {

	//is there a username?
	row := uri.conn.QueryRow("SELECT * FROM users where UserName = $1", userName)
	
	user := entity.User{}
	if row != nil {
		err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
		if err != nil {
			return user, errors.New("username  and/or password do not match")
		}

		//does the entered password match with the stred password?
	err = bcrypt.CompareHashAndPassword(user.Password,[]byte(password))
	if err!= nil{
		return user,errors.New("username and/or password do not match")
	}
		
		return user, nil
	}
	return user, errors.New("username and/or passwod do not match")


	
	
}
//GetUser ... 
func (uri *UserRepositoryImpl) GetUser(userName string) (entity.User, error) {
	// check username if exist return users
	row := uri.conn.QueryRow("SELECT * FROM users where UserName = $1", userName)
	user := entity.User{}
	if row != nil {
		err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
		if err != nil {
			return user, err
		}
		return user, nil
	}
	return user, errors.New("user not found")

}

//EditUser ... edit our user entiity
func (uri *UserRepositoryImpl) EditUser(user entity.User) error {
	_, err := uri.conn.Exec("UPDATE users SET FirstName = $1,LastName = $2,UserName = $3,Email = $4,Password= $5, Phone = $6,Image = $7 WHERE id = $8", user.FirstName, user.LastName, user.UserName, user.Email,user.Password, user.Phone, user.Image,user.UserID)
	if err != nil {
		return errors.New("Update has faild")
	}
	return nil

}

//DeleteUser ... Delete user
func (uri *UserRepositoryImpl) DeleteUser(id int) error {
	_, err := uri.conn.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return errors.New("Delete has faild")
	}
	return nil
}
