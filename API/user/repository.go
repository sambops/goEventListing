package user

import (
	"github.com/goEventListing/API/entity"
)

//EXTERNALINTERFACE(DATABASE)

//UserRepository repository(interface) spacifies User user related database operations
type UserRepository interface {
	RegisterUser(user *entity.User)(*entity.User,error)
	AuthenticateUser(userName string, password string) (*entity.User, error)
	GetUser(userName string) (*entity.User, error)
	//Logout() error
	EditUser(user *entity.User)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,error)
}
