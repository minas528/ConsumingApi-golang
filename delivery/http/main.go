package main

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/web-prog-go-sample-master/delivery/http/handler"
	"github.com/minas528/web-prog-go-sample-master/entity"
	"github.com/minas528/web-prog-go-sample-master/menu/repository"
	"github.com/minas528/web-prog-go-sample-master/menu/service"
	"html/template"
	"net/http"

	_ "github.com/jinzhu/gorm-master/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {

	dbconngorm, errgorm := gorm.Open("postgres","host=localhost port=9090 user=app_admin dbname=restaurantdb password=P@$$w0rdD2 sslmode=disable")

	if errgorm != nil{
		panic(errgorm)
	}
	defer dbconngorm.Close()

	errs := dbconngorm.CreateTable(&entity.Comment{}).GetErrors()

	//
	//dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$w0rdD2@localhost:9090/restaurantdb?sslmode=disable")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer dbconn.Close()
	//
	//if err := dbconn.Ping(); err != nil {
	//	panic(err)
	//}

	tmpl := template.Must(template.ParseGlob("ui/templates/*.html"))

	//categoryRepo := repository.NewCategoryRepositoryImpl(dbconn)
	//categoryServ := service.NewCategoryServiceImpl(categoryRepo)

	categorRepoGorm := repository.CategoryRepositoryImple2{dbconngorm}
	categoryServGorm := service.NewCategoryServiceImpl(&categorRepoGorm)

	//roleRepo := repository.NewRoleRepository(dbconn)
	//roleServ := service.NewRoleServiceImpl(roleRepo)
	//
	//userRepo := repository.NewUserRepository(dbconn)
	//userServ := service.NewUserServiceImple(userRepo)
	//
	adminCatgHandler := handler.NewAdminCategoryHandler(tmpl, categoryServGorm)
	//menuHandler := handler.NewMenuHandler(tmpl, categoryServ)
	//roleHandler := handler.NewAdminRoleHandler(tmpl,roleServ)
	//userHandler := handler.NewAdminUserHandler(tmpl,userServ)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//http.HandleFunc("/", menuHandler.Index)
	//http.HandleFunc("/about", menuHandler.About)
	//http.HandleFunc("/contact", menuHandler.Contact)
	//http.HandleFunc("/menu", menuHandler.Menu)
	//http.HandleFunc("/admin", menuHandler.Admin)
	//
	http.HandleFunc("/admin/categories", adminCatgHandler.AdminCategories)
	http.HandleFunc("/admin/categories/new", adminCatgHandler.AdminCategoriesNew)
	http.HandleFunc("/admin/categories/update", adminCatgHandler.AdminCategoriesUpdate)
	http.HandleFunc("/admin/categories/delete", adminCatgHandler.AdminCategoriesDelete)
	//
	//http.HandleFunc("/admin/roles/new",roleHandler.AdminRolesNew)
	////log.Println("on home page")
	//http.HandleFunc("/admin/users",userHandler.AdminUser)
	//http.HandleFunc("/admin/users/new",userHandler.AdminUsersNew)
	//http.HandleFunc("/admin/users/update",userHandler.AdminUsersUpdate)
	//http.HandleFunc("/admin/users/delete",userHandler.AdminUsersDelete)


	http.ListenAndServe(":8181", nil)
}
