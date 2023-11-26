package pegawaicontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	pegawaiservice "github.com/fzialam/workAway/service/pegawai_service"
	"github.com/gorilla/mux"
)

type PegawaiControllerImpl struct {
	PegawaiService pegawaiservice.PegawaiService
}

func NewPegawaiController(pegawaiService pegawaiservice.PegawaiService) PegawaiController {
	return &PegawaiControllerImpl{
		PegawaiService: pegawaiService,
	}
}

// Index implements PegawaiController.
func (pc *PegawaiControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	allUserResponse := pc.PegawaiService.GetAllUserId(r.Context())
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   allUserResponse,
	}
	temp, err := template.ParseFiles("./view/permohonan.html")
	helper.PanicIfError(err)
	temp.Execute(w, response)
}

// CreatePermohonan implements PegawaiController.
func (pc *PegawaiControllerImpl) CreatePermohonan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	permohonanRequest := permohonanreqres.PermohonanRequest{
		UserPemohonId: userId,
	}
	helper.ReadFromRequestBody(r, &permohonanRequest)

	permohonanResponse := pc.PegawaiService.CreatePermohonan(r.Context(), permohonanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   permohonanResponse,
	}
	helper.WriteToResponseBody(w, response)
}

// Presensi implements PegawaiController.
func (pc *PegawaiControllerImpl) Presensi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	presensiRequest := presensireqres.PresensiFotoRequest{
		UserId: userId,
	}
	helper.ReadFromRequestBody(r, &presensiRequest)

	presensiResponse := pc.PegawaiService.PresensiFoto(r.Context(), presensiRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   presensiResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// GetSurat implements PegawaiController.
func (pc *PegawaiControllerImpl) GetSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	getSuratRequest := presensireqres.GetSuratForPresensiRequest{
		UserId: userId,
	}
	getSuratResponse := pc.PegawaiService.GetSurat(r.Context(), getSuratRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   getSuratResponse,
	}

	helper.WriteToResponseBody(w, response)
}
