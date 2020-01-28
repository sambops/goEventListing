package user

import (
	"github.com/goEventListing/API/entity"
)

//EXTERNALINTERFACE(DATABASE)

//UserRepository repository(interface) spacifies User user related database operations
type UserRepository interface {
	RegisterUser(user *entity.User)(*entity.User,error)
	//AuthenticateUser(userName string, password string) (*entity.User, error)
	GetUser(id uint) (*entity.User, error)
	GetUsers() ([]entity.User, []error)

	GetUserByUserName(userName string) (*entity.User, error)
	//Logout() error
	EditUser(user *entity.User)(*entity.User,[]error)
	DeleteUser(id uint)(*entity.User,error)
	UserRoles(*entity.User) ([]entity.Role, []error)

	UserByEmail(email string) (*entity.User, []error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool

}

// RoleRepository speifies application user role related database operations
type RoleRepository interface {
	
	Roles() ([]entity.Role, []error)
	Role(id uint) (*entity.Role, []error)
	RoleByName(name string) (*entity.Role, []error)
	UpdateRole(role *entity.Role) (*entity.Role, []error)
	DeleteRole(id uint) (*entity.Role, []error)
	StoreRole(role *entity.Role) (*entity.Role, []error)
}

// SessionRepository specifies logged in user session related database operations
type SessionRepository interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}

