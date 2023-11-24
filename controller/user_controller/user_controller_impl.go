package usercontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	userreqes "github.com/fzialam/workAway/model/req_res/user_req_res"
	userservice "github.com/fzialam/workAway/service/user_service"
	"github.com/gorilla/mux"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// IndexL implements UserController.
func (uc *UserControllerImpl) IndexL(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./view/login.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, nil)
	helper.PanicIfError(err)
}

// IndexR implements UserController.
func (uc *UserControllerImpl) IndexR(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./view/register.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, nil)
	helper.PanicIfError(err)
}

// Login implements UserController.
func (uc *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLoginRequest := userreqes.UserLoginRequest{}
	helper.ReadFromRequestBody(r, &userLoginRequest)

	userResponse := uc.UserService.Login(r.Context(), userLoginRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Register implements UserController.
func (uc *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	userRegisterRequest := userreqes.UserRegisterRequest{}
	helper.ReadFromRequestBody(r, &userRegisterRequest)

	userResponse := uc.UserService.Register(r.Context(), userRegisterRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Logout implements UserController.
func (*UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Update implements UserController.
func (uc *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	userUpdateRequest := userreqes.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	userUpdateRequest.Id = userId

	userResponse := uc.UserService.Update(r.Context(), userUpdateRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Delete implements UserController.
func (uc *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	uc.UserService.Delete(r.Context(), userId)
	response := model.Response{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, response)
}

// FindAll implements UserController.
func (uc *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	userResponse := uc.UserService.FindAll(r.Context())
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// FindByEmail implements UserController.
func (uc *UserControllerImpl) FindByEmail(w http.ResponseWriter, r *http.Request) {
	// logical find by email
	var email string

	userResponse := uc.UserService.FindByEmail(r.Context(), email)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
	panic("unimplemented")
}

// FindByNIP implements UserController.
func (uc *UserControllerImpl) FindByNIP(w http.ResponseWriter, r *http.Request) {
	// logical find by nip
	var nip string

	userResponse := uc.UserService.FindByNIP(r.Context(), nip)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
	panic("unimplemented")
}
