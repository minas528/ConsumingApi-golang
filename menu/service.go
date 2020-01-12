package menu

import "github.com/minas528/web-prog-go-sample-master/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}

// RoleService specifies application operations different make
type RoleService interface {
	Roles() ([]entity.Role, error)
	Role(id int) (entity.Role, error)
	UpdateRole(role entity.Role) error
	DeleteRole(id int) error
	StoreRole(role entity.Role) error
}

type UserService interface {
	Users() ([]entity.Users,error)
	User(id int) (entity.Users,error)
	UpdateUser(users entity.Users) error
	DeleteUser(id int) error
	StoreUser(users entity.Users) error
}