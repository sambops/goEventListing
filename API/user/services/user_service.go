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

//GetUserByUserName ... 
func (usi *UserServiceImpl) GetUserByUserName(userName string) (*entity.User, error) {
	//check username?
	user, err := usi.userRepo.GetUserByUserName(userName)
	if err != nil {
		return user, err
	}
	return user, nil

}


//GetUser ... returns one user row with the given user name
func (usi *UserServiceImpl) GetUser(id uint) (*entity.User, error) {
	//check username?
	user, err := usi.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil

}
//GetUsers returns all stored application users
func (usi *UserServiceImpl) GetUsers() ([]entity.User, []error) {
	usrs, errs := usi.userRepo.GetUsers()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
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

// UserRoles returns list of roles a user has
func (usi *UserServiceImpl) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles, errs := usi.userRepo.UserRoles(user)

	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
// PhoneExists check if there is a user with a given phone number
func (usi *UserServiceImpl) PhoneExists(phone string) bool {
	exists := usi.userRepo.PhoneExists(phone)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (usi *UserServiceImpl) EmailExists(email string) bool {
	exists := usi.userRepo.EmailExists(email)
	return exists
}

// UserByEmail retrieves an application user by its email address
func (usi *UserServiceImpl) UserByEmail(email string) (*entity.User, []error) {
	usr, errs := usi.userRepo.UserByEmail(email)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

