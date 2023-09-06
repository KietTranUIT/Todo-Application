package controller

import (
	"net/http"
	"user-service/internal/core/port/service"
)

type UserController struct {
	Mux     *http.ServeMux
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return UserController{
		Mux:     http.NewServeMux(),
		service: service,
	}
}

func (u UserController) Router() {
	u.Mux.HandleFunc("/", Handle)
	u.Mux.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
	u.Mux.HandleFunc("/signup", SignUp(u))
	u.Mux.HandleFunc("/sendmail", SendVerifyEmail(u))
	u.Mux.HandleFunc("/verify", VerifyEmail(u))
	u.Mux.HandleFunc("/signin", SignIn(u))
	u.Mux.HandleFunc("/home", Home)
	u.Mux.HandleFunc("/createcategory", CreateCategory(u))
}
