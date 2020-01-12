package handler

import (
	"github.com/minas528/web-prog-go-sample-master/entity"
	"github.com/minas528/web-prog-go-sample-master/menu"
	"html/template"
	"net/http"
)

type AdminRoleHandler struct {
	tml *template.Template
	roleSrv menu.RoleService
}

func NewAdminRoleHandler(T *template.Template,rs menu.RoleService) *AdminRoleHandler  {
	return &AdminRoleHandler{tml:T,roleSrv:rs}
}
func (arh AdminRoleHandler) AdminRolesNew(w http.ResponseWriter,r *http.Request)  {
	//log.Println("here")

	if r.Method == http.MethodPost{
		name := r.FormValue("name")
		role := entity.Role{Name:name}
		err := arh.roleSrv.StoreRole(role)
		if err != nil{
			panic(err)
		}

		http.Redirect(w,r,"/admin/roles/new",http.StatusSeeOther)
	}
	arh.tml.ExecuteTemplate(w,"admin.roles.new.layout",nil)
}
