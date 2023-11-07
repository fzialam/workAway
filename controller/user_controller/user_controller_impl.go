package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	userreqes "github.com/fzialam/workAway/model/req_res/user_req_res"
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
func (uc *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*userAgent := r.Header.Get("User-Agent")

	// Melakukan identifikasi berdasarkan User-Agent
	if helper.UserAgentContains(userAgent, "Android") ||
		helper.UserAgentContains(userAgent, "Dart") {
		fmt.Println("Perangkat Android mengakses aplikasi Anda.")
	} else if helper.UserAgentContains(userAgent, "iPhone") {
		fmt.Println("Perangkat iPhone mengakses aplikasi Anda.")
	} else if helper.UserAgentContains(userAgent, "Chrome") {
		fmt.Println("Perangkat tidak dikenali.", userAgent)
	} else {
		fmt.Println("Perangkat tidak dikenali.", userAgent)
	}*/
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
func (uc *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

// Update implements UserController.
func (uc *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdateRequest := userreqes.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userId := p.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := uc.UserService.Update(r.Context(), userUpdateRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Delete implements UserController.
func (uc *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Kurang Logical Delete Account
	userId := p.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	uc.UserService.Delete(r.Context(), id)
	response := model.Response{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, response)
}

// FindAll implements UserController.
func (uc *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userResponse := uc.UserService.FindAll(r.Context())
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// FindByEmail implements UserController.
func (uc *UserControllerImpl) FindByEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
func (uc *UserControllerImpl) FindByNIP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
