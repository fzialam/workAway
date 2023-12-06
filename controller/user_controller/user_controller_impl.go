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
	temp, err := template.ParseFiles("view/login.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, nil)
	helper.PanicIfError(err)
}

// IndexR implements UserController.
func (uc *UserControllerImpl) IndexR(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/register.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, nil)
	helper.PanicIfError(err)
}

// Login implements UserController.
func (uc *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {

	userLoginRequest := userreqes.UserLoginRequest{}
	isMobbile := r.URL.Query().Get("m")
	if isMobbile != "" {
		helper.ReadFromRequestBody(r, &userLoginRequest)

		userResponse, err := uc.UserService.Login(r.Context(), userLoginRequest)
		helper.PanicIfError(err)

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt-workaway",
			Value:    userResponse.Token,
			HttpOnly: true,
			Path:     "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "userId",
			Value:    strconv.Itoa(userResponse.Id),
			HttpOnly: true,
			Path:     "/",
		})

		helper.WriteToResponseBody(w, userResponse)
	} else {
		err := r.ParseForm()
		helper.PanicIfError(err)
		var data map[string]interface{}

		userLoginRequest = userreqes.UserLoginRequest{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}
		userResponse, err := uc.UserService.Login(r.Context(), userLoginRequest)
		if err != nil {
			data = map[string]interface{}{
				"success": 0,
				"message": err,
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt-workaway",
			Value:    userResponse.Token,
			HttpOnly: true,
			Path:     "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "userId",
			Value:    strconv.Itoa(userResponse.Id),
			HttpOnly: true,
			Path:     "/",
		})

		if userResponse.Rank == 0 {
			data = map[string]interface{}{
				"success":  1,
				"redirect": "/wp/" + strconv.Itoa(userResponse.Id) + "/home",
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
		} else if userResponse.Rank == 1 {
			data = map[string]interface{}{
				"success":  1,
				"redirect": "/wpp/home",
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
		} else if userResponse.Rank == 2 {
			data = map[string]interface{}{
				"success":  1,
				"redirect": "/wt/home",
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
		} else if userResponse.Rank == 3 {
			data = map[string]interface{}{
				"success":  1,
				"redirect": "/wk/home",
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
		} else if userResponse.Rank == 4 {
			data = map[string]interface{}{
				"success":  1,
				"redirect": "/wa/home",
			}
			temp, err := template.ParseFiles("view/login.html")
			helper.PanicIfError(err)

			err = temp.Execute(w, data)
			helper.PanicIfError(err)
		}
	}

}

// Register implements UserController.
func (uc *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.PanicIfError(err)

	s := r.Form.Get("gender")

	gender, err := strconv.Atoi(s)
	helper.PanicIfError(err)

	userRegisterRequest := userreqes.UserRegisterRequest{
		NIP:      r.Form.Get("nip"),
		NIK:      r.Form.Get("nik"),
		NPWP:     r.Form.Get("npwp"),
		Name:     r.Form.Get("fullname"),
		NoTelp:   r.Form.Get("noTelp"),
		TglLahir: r.Form.Get("lahir"),
		Status:   "1",
		Gender:   gender,
		Alamat:   r.Form.Get("alamat"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	userResponse := uc.UserService.Register(r.Context(), userRegisterRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	t, err := template.ParseFiles("view/success.html")
	helper.PanicIfError(err)

	t.Execute(w, response)
}

// Logout implements UserController.
func (uc *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "jwt-workaway",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "userId",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
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

// ChangePassword implements UserController.
func (uc *UserControllerImpl) ChangePassword(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	changePassReq := userreqes.ChangePasswordReq{
		Id: id,
	}
	helper.ReadFromRequestBody(r, &changePassReq)

	userResponse := uc.UserService.ChangePassword(r.Context(), changePassReq)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// UpdateProfile implements UserController.
func (uc *UserControllerImpl) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	userUpdateRequest := userreqes.UserUpdateRequest{
		Id: id,
	}

	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userResponse := uc.UserService.UpdateProfile(r.Context(), userUpdateRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// ChangeImage implements UserController.
func (uc *UserControllerImpl) ChangeImage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	changeImageRequest := userreqes.ChangeImageReq{
		Id: id,
	}

	helper.ReadFromRequestBody(r, &changeImageRequest)

	userResponse := uc.UserService.ChangeImage(r.Context(), changeImageRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, response)
}
