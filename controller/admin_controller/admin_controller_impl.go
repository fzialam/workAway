package admincontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	userreqres "github.com/fzialam/workAway/model/req_res/user_req_res"
	adminservice "github.com/fzialam/workAway/service/admin_service"
	"github.com/gorilla/mux"
)

type AdminControllerImpl struct {
	AdminService adminservice.AdminService
}

func NewAdminController(adminService adminservice.AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}

// Index implements AdminController.
func (ac *AdminControllerImpl) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"menu": "home",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// Permohonan implements AdminController.
func (ac *AdminControllerImpl) Permohonan(w http.ResponseWriter, r *http.Request) {
	permohonan, err := ac.AdminService.Permohonan(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"permohonan": permohonan,
		"menu":       "permohonan",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// Penugasan implements AdminController.
func (ac *AdminControllerImpl) Penugasan(w http.ResponseWriter, r *http.Request) {
	penugasan, err := ac.AdminService.Penugasan(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"penugasan": penugasan,
		"menu":      "penugasan",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// LapAKK implements AdminController.
func (ac *AdminControllerImpl) LapAKK(w http.ResponseWriter, r *http.Request) {
	laporanAKK, err := ac.AdminService.LapAKK(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"laporan": laporanAKK,
		"menu":    "laporanAKK",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// LapAGG implements AdminController.
func (ac *AdminControllerImpl) LapAGG(w http.ResponseWriter, r *http.Request) {
	laporanAGG, err := ac.AdminService.LapAGG(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"laporan": laporanAGG,
		"menu":    "laporanAGG",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// UserGET implements AdminController.
func (ac *AdminControllerImpl) UserGET(w http.ResponseWriter, r *http.Request) {
	userResponses, err := ac.AdminService.UserGET(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"users": userResponses,
		"menu":  "users",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// UserGETById implements AdminController.
func (ac *AdminControllerImpl) UserGETById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["userId"]

	id, err := strconv.Atoi(idS)
	helper.PanicIfError(err)

	userById, err := ac.AdminService.UserGETById(r.Context(), id)
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"user": userById,
		"menu": "userById",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// UserPOST implements AdminController.
func (ac *AdminControllerImpl) UserPOST(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idS := vars["userId"]

	id, err := strconv.Atoi(idS)
	helper.PanicIfError(err)

	req := userreqres.RankChangeRequest{
		Id: id,
	}

	helper.ReadFromRequestBody(r, &req)

	err = ac.AdminService.UserPOST(r.Context(), req)
	helper.PanicIfError(err)

	response := model.Response{
		Code:   200,
		Status: "Ok",
		Data:   req,
	}

	helper.WriteToResponseBody(w, response)

}

// Profile implements AdminController.
func (ac *AdminControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	userResponse := ac.AdminService.Profile(r.Context(), id)

	data := map[string]interface{}{
		"user": userResponse,
		"menu": "profile",
	}
	temp, err := template.ParseFiles("view/admin.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
