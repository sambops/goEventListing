package repository

// import (
// 	"github.com/goEventListing/API/entity"
// 	"github.com/goEventListing/API/user"
// 	"github.com/jinzhu/gorm"
// )

// // MockRoleRepo implements the menu.RoleRepository interface
// type MockRoleRepo struct {
// 	conn *gorm.DB
// }

// // NewMockRoleRepo returns a new a new object of RoleGormRepo
// func NewMockRoleRepo(db *gorm.DB) user.RoleRepository {
// 	return &MockRoleRepo{conn: db}
// }

// // Roles returns all user roles stored in the database
// func (roleRepo *MockRoleRepo) Roles() ([]entity.Role, []error) {

// 	rls := []entity.Role{entity.MockRole}
// 	return rls, nil
// }

// // Role retrieves a role by its id from the database
// func (roleRepo *MockRoleRepo) Role(id uint) (*entity.Role, []error) {
// 	rl := entity.MockRole
// 	if id == 1 {
// 		return &rl, nil
// 	}
// 	return nil, nil
// }

// // RoleByName retrieves a role by its name from the database
// func (roleRepo *MockRoleRepo) RoleByName(name string) (*entity.Role, []error) {
// 	rl := entity.MockRole
// 	if name == "teme" {
// 		return &rl, nil
// 	}
// 	return nil, nil
// }

// // UpdateRole updates a given user role in the database
// func (roleRepo *MockRoleRepo) UpdateRole(role *entity.Role) (*entity.Role, []error) {
// 	usr := entity.MockRole
// 	return &usr, nil

// }

// // DeleteRole deletes a given user role from the database
// func (roleRepo *MockRoleRepo) DeleteRole(id uint) (*entity.Role, []error) {
// 	rl := entity.MockRole

// 	if id == 1 {
// 		return &rl, nil
// 	}
// 	return &rl, nil
// }

// // StoreRole stores a given user role in the database
// func (roleRepo *MockRoleRepo) StoreRole(role *entity.Role) (*entity.Role, []error) {
// 	rl := role
// 	return rl, nil

// }
