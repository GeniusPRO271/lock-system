package role

import (
	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"gorm.io/gorm"
)

type RoleService interface {
	CreateRole(role *model.Role) (err error)
	GetRoles(Role *[]model.Role) (err error)
	GetRole(Role *model.Role, id int) (err error)
	UpdateRole(Role *model.Role) (err error)
}

type RoleServiceImpl struct {
	// You may include dependencies here, such as a database connection
	// db *Database
	Db *gorm.DB
}

// Create a role
func (s RoleServiceImpl) CreateRole(Role *model.Role) (err error) {
	err = s.Db.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func (s RoleServiceImpl) GetRoles(Role *[]model.Role) (err error) {
	err = s.Db.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func (s RoleServiceImpl) GetRole(Role *model.Role, id int) (err error) {
	err = s.Db.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func (s RoleServiceImpl) UpdateRole(Role *model.Role) (err error) {
	s.Db.Save(Role)
	return nil
}
