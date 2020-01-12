package service

import (
	"errors"
	"github.com/minas528/web-prog-go-sample-master/entity"
	"github.com/minas528/web-prog-go-sample-master/menu"
)

type RoleServiceImpl struct {
	roleRepo menu.RoleService
}

func NewRoleServiceImpl(RoleRepo menu.RoleService) *RoleServiceImpl {
	return &RoleServiceImpl{roleRepo:RoleRepo}
}

func (rr *RoleServiceImpl) StoreRole(role entity.Role) error {
	err := rr.roleRepo.StoreRole(role)
	if err != nil{
		return errors.New("Storing roles has failed")
	}
	return nil
}
func (rr *RoleServiceImpl) Roles() ([]entity.Role, error){
	return []entity.Role{},nil
}
func (rr *RoleServiceImpl)  Role(id int) (entity.Role, error){
	return entity.Role{},nil
}
func (rr *RoleServiceImpl) UpdateRole(role entity.Role) error{
	return nil
}
func (rr *RoleServiceImpl) DeleteRole(id int)  error{
	return nil
}