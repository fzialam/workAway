package permohonancontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	permohonanservice "github.com/fzialam/workAway/service/permohonan_service"
	"github.com/gorilla/mux"
)

type PermohonanControllerImpl struct {
	PermohonanService permohonanservice.PermohonanService
}

func NewPermohonanController(permohonanService permohonanservice.PermohonanService) PermohonanController {
	return &PermohonanControllerImpl{
		PermohonanService: permohonanService,
	}
}

// Index implements PermohonanController.
func (pc *PermohonanControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	allUserResponse := pc.PermohonanService.GetAllUserId(r.Context())
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   allUserResponse,
	}
	temp, err := template.ParseFiles("./view/permohonan.html")
	helper.PanicIfError(err)
	temp.Execute(w, response)
}

// CreatePermohonan implements PermohonanController.
func (pc *PermohonanControllerImpl) CreatePermohonan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	permohonanRequest := permohonanreqres.PermohonanRequest{
		UserPemohonId: userId,
	}
	helper.ReadFromRequestBody(r, &permohonanRequest)

	permohonanResponse := pc.PermohonanService.CreatePermohonan(r.Context(), permohonanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   permohonanResponse,
	}
	helper.WriteToResponseBody(w, response)
}
