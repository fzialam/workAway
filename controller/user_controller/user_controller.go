package usercontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindByNIP(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
