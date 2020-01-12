package repository

import (
	"database/sql"
	"errors"
	"github.com/minas528/web-prog-go-sample-master/entity"
)

type RoleRepositoryImpl struct {
	conn *sql.DB
}

func NewRoleRepository(conn *sql.DB) *RoleRepositoryImpl{
	return &RoleRepositoryImpl{conn:conn}
}

func (rr *RoleRepositoryImpl) StoreRole(role entity.Role) error  {
	query := "INSERT INTO roles (name) values ($1)"
	_,err := rr.conn.Exec(query,role.Name)

	if err != nil {
		errors.New("Insert has Failed!")
	}
	return nil
}
func (rr *RoleRepositoryImpl) Roles() ([]entity.Role, error) {
	return []entity.Role{} ,nil
}
func (rr *RoleRepositoryImpl) Role(id int) (entity.Role, error) {
	return entity.Role{}, nil
}
func (rr *RoleRepositoryImpl) UpdateRole(role entity.Role) error {
	return nil
}
func (rr *RoleRepositoryImpl) DeleteRole(id int) error {

	return nil
}