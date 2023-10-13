package usercontroller

import (
	"net/http"

	userservice "github.com/fzialam/workAway/service/user_service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// Login implements UserController.
func (*UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Register implements UserController.
func (*UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Update implements UserController.
func (*UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// Delete implements UserController.
func (*UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements UserController.
func (*UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindByEmail implements UserController.
func (*UserControllerImpl) FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindByNIP implements UserController.
func (*UserControllerImpl) FindByNIP(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}
