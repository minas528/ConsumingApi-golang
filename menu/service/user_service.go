package service

import (
	"github.com/minas528/web-prog-go-sample-master/entity"
	"github.com/minas528/web-prog-go-sample-master/menu"
)

type UserServiceImple struct {
	usrRepo menu.UserService
}

func NewUserServiceImple( service menu.UserService) *UserServiceImple  {
	return &UserServiceImple{usrRepo:service}
}

func (us *UserServiceImple) Users() ([]entity.Users,error) {

	users, err := us.usrRepo.Users()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserServiceImple) StoreUser(users entity.Users) error {

	err := us.usrRepo.StoreUser(users)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserServiceImple) User(id int) (entity.Users,error) {

	c, err := us.usrRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}
func (us *UserServiceImple) UpdateUser(users entity.Users) error {

	err := us.usrRepo.UpdateUser(users)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (us *UserServiceImple) DeleteUser(id int) error {

	err := us.usrRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}