package repository

import (
	"database/sql"
	"errors"
	"github.com/minas528/web-prog-go-sample-master/entity"
	"log"
)

type UserRepositoryImpl struct {
	conn *sql.DB
}

func NewUserRepository( conn *sql.DB) *UserRepositoryImpl  {
	return &UserRepositoryImpl{conn:conn}
}

func (ur *UserRepositoryImpl) Users() ([]entity.Users,error)  {
	log.Println("select")
	rows, err := ur.conn.Query("SELECT * FROM users;")
	if err != nil{
		return nil,errors.New("can not add to database")
	}
	defer rows.Close()

	usrs := []entity.Users{}
	for rows.Next()  {
		usr := entity.Users{}
		err := rows.Scan(&usr.ID,&usr.UUID,&usr.Name,&usr.Email,&usr.Phone,&usr.Password)
		if err != nil {
			return nil,err
		}
		usrs = append(usrs,usr)
	}
	return usrs,nil

}
func (ur *UserRepositoryImpl) User(id int) (entity.Users,error) {
	//log.Println("select")

	row := ur.conn.QueryRow("SELECT * FROM users WHERE id = $1", id)

	usr := entity.Users{}

	err := row.Scan(&usr.ID,&usr.UUID, &usr.Name,&usr.Email, &usr.Phone, &usr.Password)
	if err != nil {
		return usr, err
	}

	return usr, nil
}

func (ur *UserRepositoryImpl) UpdateUser(users entity.Users) error {

	_, err := ur.conn.Exec("UPDATE users SET uuid=$1 ,full_name=$2,phone=$3, password=$4,email=$5 WHERE id=$6", users.UUID, users.Name, users.Phone,users.Password,users.Email, users.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

func (ur *UserRepositoryImpl) DeleteUser(id int) error {

	_, err := ur.conn.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}
func (ur *UserRepositoryImpl) StoreUser(users entity.Users) error {

	_, err := ur.conn.Exec("INSERT INTO users (uuid,full_name,email,phone,password) values($1, $2, $3,$4,$5)", users.UUID, users.Name, users.Email,users.Phone,users.Password)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
