package presensicontroller

import (
	"log"
	"net/http"

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
	presensiRequest := presensirequestresponse.PresensiFotoRequest{}
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
func (pc *PresensiControllerImpl) GetSurat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Implement GetSurat")
}
