package menu

import "github.com/minas528/web-prog-go-sample-master/entity"

// CategoryRepository specifies menu category related database operations
type CategoryRepository interface {
	Categories() ([]entity.Category, error)
	Category(id int) (entity.Category, error)
	UpdateCategory(category entity.Category) error
	DeleteCategory(id int) error
	StoreCategory(category entity.Category) error
}

type RoleRespositry interface {
	Roles() ([]entity.Role,error)
	Role(id int) (entity.Role,error)
	UpdateRole(role entity.Role) error
	DeleteRole(id int) error
	StoreRole(role entity.Role) error
}

type UserRepository interface {
	Users() ([]entity.Users,error)
	User(id int) (entity.Users,error)
	UpdateUser(users entity.Users) error
	DeleteUser(id int) error
	StoreUser(users entity.Users) error
}
