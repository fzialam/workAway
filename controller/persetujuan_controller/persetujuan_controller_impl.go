package persetujuancontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	persetujuanreqres "github.com/fzialam/workAway/model/req_res/persetujuan_req_res"
	persetujuanservice "github.com/fzialam/workAway/service/persetujuan_service"
	"github.com/gorilla/mux"
)

type PersetujunanControllerImpl struct {
	PersetujuanService persetujuanservice.PersetujuanService
}

func NewPersetujuanController(persetujuanService persetujuanservice.PersetujuanService) PersetujunanController {
	return &PersetujunanControllerImpl{
		PersetujuanService: persetujuanService,
	}
}

// Index implements PersetujunanController.
func (ps *PersetujunanControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	response := ps.PersetujuanService.GetAllSuratTugasJOINApprovedUser(r.Context())
	data := map[string]interface{}{
		"response": response,
		"status":   r.URL.Query().Get("s"),
	}

	temp, err := template.ParseFiles("view/persetujuan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
	// helper.WriteToResponseBody(w, data)
}

// DetailSurat implements PersetujunanController.
func (ps *PersetujunanControllerImpl) DetailSurat(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := ps.PersetujuanService.GetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"response": response,
		"lenP":     len(response.Participans),
		"status":   r.URL.Query().Get("s"),
	}

	temp, err := template.ParseFiles("view/persetujuan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
	// helper.WriteToResponseBody(w, data)
}

// SetApproved implements PersetujunanController.
func (ps *PersetujunanControllerImpl) SetApproved(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["suratId"]
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	persetujuanRequest := persetujuanreqres.PersetujuanRequest{}
	helper.ReadFromRequestBody(r, &persetujuanRequest)

	persetujuanRequest.SuratTugasId = idInt
	persetujuanRequest.CreateAt = helper.TimeNowToString()

	persetujuanResponse := ps.PersetujuanService.SetApproved(r.Context(), persetujuanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   persetujuanResponse,
	}

	helper.WriteToResponseBody(w, response)
}
