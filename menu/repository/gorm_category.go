package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/web-prog-go-sample-master/entity"
	"log"
)

type CategoryRepositoryImple2 struct {
	Conn *gorm.DB
}

func (cri *CategoryRepositoryImple2) Category(id int) (entity.Category, error) {
	c := entity.Category{}
	cri.Conn.Take(&c)
	return c,nil
}

func (cri *CategoryRepositoryImple2) UpdateCategory(category entity.Category) error {
	var c entity.Category
	cri.Conn.Find(&c).Update()
}

func (cri *CategoryRepositoryImple2) DeleteCategory(id int) error {
	cri.Conn.Delete()
}

func (cri *CategoryRepositoryImple2) StoreCategory(category entity.Category) error {
	panic("implement me")
}

func NewCategoryRepositoryImple2( conn *gorm.DB) *CategoryRepositoryImple2  {
	return &CategoryRepositoryImple2{Conn:conn}
}


// Categories returns all cateogories from the database
func (cri *CategoryRepositoryImple2) Categories() ([]entity.Category, error){
	ctgs := []entity.Category{}
	cri.Conn.Find(&ctgs).GetErrors()
	return ctgs,nil
}