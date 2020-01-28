package repository

import (
	"errors"

	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/user"
	"github.com/jinzhu/gorm"
)

// MockUserGormRepo is repository implements user.UserRepository
type MockUserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates and returns new UserGormRepo object
func NewMockUserGormRepo(dbConn *gorm.DB) user.UserRepository {
	return &MockUserGormRepo{conn: dbConn}
}

//
func (uri *MockUserGormRepo) RegisterUser(user *entity.User) (*entity.User, error) {
	usr := user
	return usr, nil

}

//
func (uri *MockUserGormRepo) GetUserByUserName(userName string) (*entity.User, error) {
	usr := entity.MockUser
	if userName == "someone" {
		return &usr, nil
	}
	return nil, nil
}

//

func (uri *MockUserGormRepo) GetUser(id uint) (*entity.User, error) {
	usr := entity.MockUser
	if id == 1 {
		return &usr, nil
	}
	return nil, errors.New("Not found")
}

//
func (uri *MockUserGormRepo) GetUsers() ([]entity.User, []error) {
	usrs := []entity.User{entity.MockUser}
	return usrs, nil
}

//
func (uri *MockUserGormRepo) EditUser(user *entity.User) (*entity.User, []error) {
	usr := entity.MockUser
	return &usr, nil

}

//
func (uri *MockUserGormRepo) DeleteUser(id uint) (*entity.User, error) {
	usr := entity.MockUser
	if id != 1 {
		return nil, nil
	}

	return &usr, nil
}

//
func (uri *MockUserGormRepo) UserRoles(user *entity.User) ([]entity.Role, []error) {
	var usr []entity.User
	usr = append(usr, entity.MockUser)
	var roles []entity.Role
	rl := entity.Role{ID:1, Name:"user",Users:usr}

	roles = append(roles, rl)

	return roles, nil
}

//
func (uri *MockUserGormRepo) PhoneExists(phone string) bool {
	if phone == "0926100732" {
		return true
	}
	return false

}

//
func (uri *MockUserGormRepo) EmailExists(email string) bool {
	if email == "sura@gmail.com" {
		return true
	}
	return false
}

//
func (uri *MockUserGormRepo) UserByEmail(email string) (*entity.User, []error) {
	usr := entity.MockUser
	if email == "sura@gmail.com" {
		return &usr, nil
	}
	return nil, []error{errors.New("Not found")}
}
