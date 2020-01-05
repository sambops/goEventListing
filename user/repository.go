package user

import (
	"github.com/goEventListing/entity"
)

//EXTERNALINTERFACE(DATABASE)

//UserRepository repository(interface) spacifies User user related database operations
type UserRepository interface {
	RegisterUser(user *entity.User) error
	AuthenticateUser(userName string, password string) (entity.User, error)
	GetUser(userName string) (entity.User, error)
	//Logout() error, not important here as it is more related with cookies(http)
	EditUser(user *entity.User) []error
	DeleteUser(id int) error
}
