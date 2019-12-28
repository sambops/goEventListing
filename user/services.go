package user

import (
	"github.com/EventListing/entity"
)

//USECASE
//this is our service usescase(has (interface)abstract classes that outer layers can use)
type UserService interface {
	RegisterUser(user entity.User) error
	AuthenticateUser(userName string, password string) (entity.User, error)
	GetUser(userName string) (entity.User, error)
	//Logout() error
	EditUser(user entity.User) error
	DeleteUser(id int) error
}
