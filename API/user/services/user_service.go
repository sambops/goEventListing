package services

import (
	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/user"
)

//UserServiceImpl implements user.UserService interface
type UserServiceImpl struct {
	userRepo user.UserRepository
}

//NewUserServiceImpl ... creates an object of UserServiceImpl
func NewUserServiceImpl(UserRep user.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: UserRep}
}

//RegisterUser ... registers a new user
func (usi *UserServiceImpl) RegisterUser(user *entity.User) (*entity.User,error) {
	usr,err := usi.userRepo.RegisterUser(user)
	if err != nil {
		return usr,err
	}
	return usr,nil
}

//AuthenticateUser ... checks username and password validity
func (usi *UserServiceImpl) AuthenticateUser(userName string, password string) (*entity.User, error) {
	user, err := usi.userRepo.AuthenticateUser(userName, password)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetUser ... returns one user row with the given user name
func (usi *UserServiceImpl) GetUser(userName string) (*entity.User, error) {
	//check username?
	user, err := usi.userRepo.GetUser(userName)
	if err != nil {
		return user, err
	}
	return user, nil

}

//EditUser ... edit existing user data(profile)
func (usi *UserServiceImpl) EditUser(user *entity.User)(*entity.User,[]error) {
	urs,err := usi.userRepo.EditUser(user)
	if err != nil {
		return urs,err
	}
	return urs,nil
}

//DeleteUser ... delete existing user with the given id
func (usi *UserServiceImpl) DeleteUser(id uint)(*entity.User,error) {
	urs,err := usi.userRepo.DeleteUser(id)
	if err != nil {
		return urs,nil
	}
	return urs,nil
}
