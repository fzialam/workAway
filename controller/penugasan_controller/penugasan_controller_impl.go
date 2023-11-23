package penugasancontroller

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	penugasanservice "github.com/fzialam/workAway/service/penugasan_service"
	"github.com/gorilla/mux"
)

type PenugasanControllerImpl struct {
	PenugasanService penugasanservice.PenugasanService
}

func NewPenugasanController(penugasanService penugasanservice.PenugasanService) PenugasanController {
	return &PenugasanControllerImpl{
		PenugasanService: penugasanService,
	}
}

// Index implements PenugasanController.
func (pc *PenugasanControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	create := r.URL.Query().Get("c")

	result := strings.Compare(create, "true")

	if result == 0 {
		allUserResponse := pc.PenugasanService.GetAllUserId(r.Context())
		data := map[string]interface{}{
			"user": allUserResponse,
		}

		temp, err := template.ParseFiles("./view/penugasan-form.html")
		helper.PanicIfError(err)

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	} else {
		allSuratResponse := pc.PenugasanService.GetAllSuratTugasJOINApprovedUser(r.Context())
		data := map[string]interface{}{
			"surat":  allSuratResponse,
			"status": r.URL.Query().Get("v"),
		}

		temp, err := template.ParseFiles("./view/penugasan.html")
		helper.PanicIfError(err)

		temp.Funcs(template.FuncMap{"index": helper.AddIndex})

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	}
}

// CreatePermohonan implements PenugasanController.
func (pc *PenugasanControllerImpl) CreatePenugasan(w http.ResponseWriter, r *http.Request) {
	permohonanRequest := penugasanreqres.PenugasanRequest{}
	helper.ReadFromRequestBody(r, &permohonanRequest)

	permohonanResponse := pc.PenugasanService.CreatePenugasan(r.Context(), permohonanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   permohonanResponse,
	}
	helper.WriteToResponseBody(w, response)
}

// SetApproved implements PenugasanController.
func (pc *PenugasanControllerImpl) SetApproved(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["suratId"]
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	izinRequest := izinreqres.IzinRequest{}
	helper.ReadFromRequestBody(r, &izinRequest)

	izinRequest.SuratTugasId = idInt
	izinRequest.StatusTTD = "0"

	persetujuanResponse := pc.PenugasanService.SetApproved(r.Context(), izinRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   persetujuanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// DetailSurat implements PenugasanController.
func (pc *PenugasanControllerImpl) DetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := pc.PenugasanService.GetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"surat":  response,
		"status": r.URL.Query().Get("v"),
	}

	temp, err := template.ParseFiles("./view/penugasan.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
