package presensicontroller

import (
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	presensiservice "github.com/fzialam/workAway/service/presensi_service"
	"github.com/gorilla/mux"
)

type PresensiControllerImpl struct {
	PresensiService presensiservice.PresensiService
}

func NewPresensiController(presensiService presensiservice.PresensiService) PresensiController {
	return &PresensiControllerImpl{
		PresensiService: presensiService,
	}
}

// Presensi implements PresensiController.
func (pc *PresensiControllerImpl) Presensi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	presensiRequest := presensireqres.PresensiFotoRequest{
		UserId: userId,
	}
	helper.ReadFromRequestBody(r, &presensiRequest)

	presensiResponse := pc.PresensiService.PresensiFoto(r.Context(), presensiRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   presensiResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// GetSurat implements PresensiController.
func (pc *PresensiControllerImpl) GetSuratForPresensi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	getSuratRequest := presensireqres.GetSuratForPresensiRequest{
		UserId: userId,
	}
	getSuratResponse := pc.PresensiService.GetSurat(r.Context(), getSuratRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   getSuratResponse,
	}

	helper.WriteToResponseBody(w, response)
}
