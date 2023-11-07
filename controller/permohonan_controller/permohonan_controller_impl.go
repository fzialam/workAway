package permohonancontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	permohonanservice "github.com/fzialam/workAway/service/permohonan_service"
	"github.com/julienschmidt/httprouter"
)

type PermohonanControllerImpl struct {
	PermohonanService permohonanservice.PermohonanService
}

// CreatePermohonan implements PermohonanController.
func NewPermohonanController(permohonanService permohonanservice.PermohonanService) PermohonanController {
	return &PermohonanControllerImpl{
		PermohonanService: permohonanService,
	}
}

type Option struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (pc *PermohonanControllerImpl) CreatePermohonan(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId, err := strconv.Atoi(p.ByName("userId"))
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

// Index implements PermohonanController.
func (pc *PermohonanControllerImpl) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
