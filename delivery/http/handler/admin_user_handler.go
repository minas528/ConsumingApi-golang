package handler

import (
	"github.com/minas528/web-prog-go-sample-master/entity"
	"github.com/minas528/web-prog-go-sample-master/menu"
	"html/template"
	"net/http"
	"strconv"
)

type AdminUserHandler struct {
	templ *template.Template
	userServ menu.UserService
}

func NewAdminUserHandler(template *template.Template, service menu.UserService) *AdminUserHandler  {
	//log.Println("creating user handler")
	return &AdminUserHandler{template,service}
}
func (auh *AdminUserHandler) AdminUser(w http.ResponseWriter, r *http.Request) {
	//log.Println("handler.user")
	users, err := auh.userServ.Users()
	if err != nil {
		panic(err)
	}
	auh.templ.ExecuteTemplate(w, "admin.users.layout", users)
}
func (auh *AdminUserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {
	//log.Println("there")


	if r.Method == http.MethodPost {

		usr := entity.Users{}
		usr.Name = r.FormValue("name")
		usr.Email = r.FormValue("email")
		usr.UUID = r.FormValue("uuid")
		usr.Phone = r.FormValue("phone")
		usr.Password = r.FormValue("password")


		err := auh.userServ.StoreUser(usr)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	}
	auh.templ.ExecuteTemplate(w, "admin.users.new.layout", nil)
}
func (auh *AdminUserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := auh.userServ.User(id)

		if err != nil {
			panic(err)
		}

		auh.templ.ExecuteTemplate(w, "admin.users.update.layout", cat)

	} else if r.Method == http.MethodPost {

		usr := entity.Users{}
		usr.ID, _ = strconv.Atoi(r.FormValue("id"))
		usr.Name = r.FormValue("name")
		usr.UUID = r.FormValue("uuid")
		usr.Email = r.FormValue("email")
		usr.Phone = r.FormValue("phone")
		usr.Password = r.FormValue("password")


		err := auh.userServ.UpdateUser(usr)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}

}

func (auh *AdminUserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = auh.userServ.DeleteUser(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}


