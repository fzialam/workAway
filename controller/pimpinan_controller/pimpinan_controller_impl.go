package pimpinancontroller

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	penugasanreqres "github.com/fzialam/workAway/model/req_res/penugasan_req_res"
	pimpinanreqres "github.com/fzialam/workAway/model/req_res/pimpinan_req_res"
	pimpinanservice "github.com/fzialam/workAway/service/pimpinan_service"
	"github.com/gorilla/mux"
)

type PimpinanControllerImpl struct {
	PimpinanService pimpinanservice.PimpinanService
}

func NewPimpinanController(pimpinanService pimpinanservice.PimpinanService) PimpinanController {
	return &PimpinanControllerImpl{
		PimpinanService: pimpinanService,
	}
}

// IndexPermohonan implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexPermohonan(w http.ResponseWriter, r *http.Request) {
	response := pc.PimpinanService.PermohonanGetAllSuratTugasJOINApprovedUser(r.Context())
	data := map[string]interface{}{
		"response": response,
		"status":   r.URL.Query().Get("s"),
	}

	temp, err := template.ParseFiles("view/persetujuan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// Index implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexSPPD(w http.ResponseWriter, r *http.Request) {
	create := r.URL.Query().Get("c")

	result := strings.Compare(create, "true")

	if result == 0 {
		allUserResponse := pc.PimpinanService.GetAllUserId(r.Context())
		data := map[string]interface{}{
			"user":   allUserResponse,
			"create": "true",
		}

		temp, err := template.ParseFiles("./view/pimpinan.html")
		helper.PanicIfError(err)

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	} else {
		allSuratResponse := pc.PimpinanService.SPPDGetAllSuratTugasJOINApprovedUser(r.Context())
		data := map[string]interface{}{
			"surat":  allSuratResponse,
			"status": r.URL.Query().Get("v"),
		}

		temp, err := template.ParseFiles("./view/pimpinan.html")
		helper.PanicIfError(err)

		temp.Funcs(template.FuncMap{"index": helper.AddIndex})

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	}
}

// CreatePenugasan implements PimpinanController.
func (pc *PimpinanControllerImpl) CreatePenugasan(w http.ResponseWriter, r *http.Request) {
	penugasanRequest := penugasanreqres.PenugasanRequest{}
	helper.ReadFromRequestBody(r, &penugasanRequest)

	penugasanResponse := pc.PimpinanService.CreatePenugasan(r.Context(), penugasanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   penugasanResponse,
	}
	helper.WriteToResponseBody(w, response)
}

// SPPDDetailSurat implements PimpinanController.
func (pc *PimpinanControllerImpl) SPPDDetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := pc.PimpinanService.SPPDGetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"surat":  response,
		"status": r.URL.Query().Get("v"),
	}

	temp, err := template.ParseFiles("./view/pimpinan.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// SPPDSetApproved implements PimpinanController.
func (pc *PimpinanControllerImpl) SPPDSetApproved(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["suratId"]
	suratId, _ := strconv.Atoi(idS)

	sppdAprroved := pimpinanreqres.UploadSPPDRequest{}
	helper.ReadFromRequestBody(r, &sppdAprroved)

	sppdAprroved.SuratTugasId = suratId

	persetujuanResponse := pc.PimpinanService.SPPDSetApproved(r.Context(), sppdAprroved)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   persetujuanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// PermohonanDetailSurat implements PimpinanController.
func (pc *PimpinanControllerImpl) PermohonanDetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := pc.PimpinanService.PermohonanGetSuratTugasById(r.Context(), idSurat)

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
}

// PermohonanSetApproved implements PimpinanController.
func (pc *PimpinanControllerImpl) PermohonanSetApproved(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["suratId"]
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	izinRequest := izinreqres.IzinRequest{}
	helper.ReadFromRequestBody(r, &izinRequest)

	izinRequest.SuratTugasId = idInt
	izinRequest.StatusTTD = "0"

	persetujuanResponse := pc.PimpinanService.PermohonanSetApproved(r.Context(), izinRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   persetujuanResponse,
	}

	helper.WriteToResponseBody(w, response)
}
