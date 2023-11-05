package presensicontroller

import (
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	presensirequestresponse "github.com/fzialam/workAway/model/presensi_request_response"
	presensiservice "github.com/fzialam/workAway/service/presensi_service"
	"github.com/julienschmidt/httprouter"
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
func (pc *PresensiControllerImpl) Presensi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := strconv.Atoi(p.ByName("userId"))
	helper.PanicIfError(err)
	presensiRequest := presensirequestresponse.PresensiFotoRequest{
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
func (pc *PresensiControllerImpl) GetSuratForPresensi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := strconv.Atoi(p.ByName("userId"))
	helper.PanicIfError(err)
	getSuratRequest := presensirequestresponse.GetSuratForPresensiRequest{
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
