package user

import (
	"github.com/goEventListing/API/entity"
)

//USECASE

//UserService ... this is our service usescase(has (interface)abstract classes that outer layers can use)
type UserService interface {
	RegisterUser(user *entity.User)(*entity.User,error)
	AuthenticateUser(userName string, password string) (*entity.User, error)
	GetUser(id uint) (*entity.User, error)
	GetUserByUserName(userName string) (*entity.User, error)
	//Logout() error
	EditUser(user *entity.User)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,error)
}
