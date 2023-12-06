package usercontroller

import (
	"net/http"
)

type UserController interface {
	IndexL(w http.ResponseWriter, r *http.Request)
	IndexR(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByNIP(writer http.ResponseWriter, request *http.Request)
	FindByEmail(writer http.ResponseWriter, request *http.Request)
	FindAll(writer http.ResponseWriter, request *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ChangeImage(w http.ResponseWriter, r *http.Request)
}
