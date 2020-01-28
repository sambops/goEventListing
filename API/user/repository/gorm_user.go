package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
	"errors"

	"github.com/goEventListing/API/entity"
)

//UserRepositoryImpl ... implements the User.UserRepository interface
type UserRepositoryImpl struct {
	conn *gorm.DB
}

//NewUserRepositoryImpl will create an object of  UserReposiotryImpl
func NewUserRepositoryImpl(Conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}

//Methods with pointer receivers can modify the value to which the receiver points.
//Since methods often need to modify their receiver, pointer receivers are more common than value receivers.



//RegisterUser ... this is a method to register our users int to the user table
func (uri *UserRepositoryImpl) RegisterUser(user *entity.User)(*entity.User ,error) {
	userr := user
	//username taken?
	_,err := uri.conn.Raw("SELECT * FROM users WHERE user_name = ?",user.UserName).Rows() 
	if err != nil{
		return nil,errors.New("user name already taken try other")
	}
	errs := uri.conn.Create(userr).GetErrors()

	if len(errs) > 0 {
		fmt.Println("check me")
		return nil, errors.New("insertion has failed")
	}
	return userr,nil

}


//GetUserByUserName ... 
func (uri *UserRepositoryImpl) GetUserByUserName(userName string) (*entity.User, error) {
	user:=entity.User{}


	//check username if exist reutrn users
	rows,err := uri.conn.Raw("SELECT * FROM users WHERE user_name = ?",userName).Rows()
	if rows != nil{
		for rows.Next(){
			uri.conn.ScanRows(rows,&user)
		}
		if err != nil{
			return &user,err
		}
		return &user,nil
	}
	return &user,errors.New("user not found")
}
//GetUser ... 
func (uri *UserRepositoryImpl) GetUser(id uint) (*entity.User, error) {
	user:=entity.User{}


	//check username if exist reutrn users
	// rows,err := uri.conn.Raw("SELECT * FROM users WHERE id = ?",id).Rows()
	// if rows != nil{
	// 	for rows.Next(){
	// 		uri.conn.ScanRows(rows,&user)
	// 	}
	// 	if err != nil{
	// 		return &user,err
	// 	}
	// 	return &user,nil
	// }

	errs := uri.conn.First(&user, id).GetErrors()

	if len(errs) > 0{
		return nil, errors.New("user not found")
	}

	return &user,nil



	// // check username if exist return users
	// row := uri.conn.QueryRow("SELECT * FROM users where username = $1", userName)
	// user := entity.User{}
	// if row != nil {
	// 	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.UserName, &user.Email,&user.Password, &user.Phone, &user.Image)
	// 	if err != nil {
	// 		return user, err
	// 	}
	// 	return user, nil
	// }
	// return user, errors.New("user not found")

}

//GetUsers return all users from the database
func (uri *UserRepositoryImpl) GetUsers() ([]entity.User, []error) {
	users := []entity.User{}
	errs := uri.conn.Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return users, errs
}


//EditUser ... edit our user entiity
func (uri *UserRepositoryImpl) EditUser(user *entity.User)(*entity.User ,[]error) {
	usr:= user
	errs :=uri.conn.Save(usr).GetErrors()

	if len(errs)>0{
		return nil,errs
	}
	return usr,nil

		// _, err := uri.conn.Exec("UPDATE users SET first_name = $1,last_name = $2,username = $3,email = $4,password= $5, phone = $6,image = $7 WHERE id = $8", user.FirstName, user.LastName, user.UserName, user.Email,user.Password, user.Phone, user.Image,user.UserID)
		// if err != nil {
		// 	return errors.New("Update has faild")
		// }
		// return nil

}

//DeleteUser ... Delete user
func (uri *UserRepositoryImpl) DeleteUser(id uint) (*entity.User,error) {
	user := entity.User{}
	rows,err:= uri.conn.Raw("DELETE FROM users WHERE id = ?",id).Rows()
	if rows != nil{
	for rows.Next(){
		uri.conn.ScanRows(rows,&user)
	}
	if err != nil{
		return &user,err
	}
	return &user,nil
		
	}
	return &user,errors.New("user not found")


		// _, err := uri.conn.Exec("DELETE FROM users WHERE id = $1", id)
		// if err != nil {
		// 	return errors.New("Delete has faild")
		// }
		// return nil
}

// UserRoles returns list of application roles that a given user has
func (uri *UserRepositoryImpl) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles := []entity.Role{}

	errs := uri.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

// PhoneExists check if a given phone number is found
func (uri *UserRepositoryImpl) PhoneExists(phone string) bool {
	user := entity.User{}
	errs := uri.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// EmailExists check if a given email is found
func (uri *UserRepositoryImpl) EmailExists(email string) bool {
	user := entity.User{}
	errs := uri.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// UserByEmail retrieves a user by its email address from the database
func (uri *UserRepositoryImpl) UserByEmail(email string) (*entity.User, []error) {
	user := entity.User{}
	errs := uri.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

