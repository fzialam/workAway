package pimpinancontroller

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	izinreqres "github.com/fzialam/workAway/model/req_res/izin_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
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

// Index implements PimpinanController.
func (pc *PimpinanControllerImpl) Index(w http.ResponseWriter, r *http.Request) {
	index, err := pc.PimpinanService.Index(r.Context())
	helper.PanicIfError(err)
	data := map[string]interface{}{
		"index": index,
		"menu":  "home",
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// IndexPenugasan implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexPenugasan(w http.ResponseWriter, r *http.Request) {
	index, err := pc.PimpinanService.IndexPenugasan(r.Context())
	helper.PanicIfError(err)
	data := map[string]interface{}{
		"surats": index,
		"menu":   "penugasan",
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// IndexPermohonan implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexPermohonan(w http.ResponseWriter, r *http.Request) {
	permohonan := pc.PimpinanService.PermohonanGetAllSuratTugasJOINApprovedUser(r.Context())
	data := map[string]interface{}{
		"permohonan": permohonan,
		"menu":       "permohonan",
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
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

// PermohonanDetailSurat implements PimpinanController.
func (pc *PimpinanControllerImpl) PermohonanDetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	permohonan := pc.PimpinanService.PermohonanGetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"permohonan": permohonan,
		"lenP":       len(permohonan.Participans),
		"menu":       "permohonanView",
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
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
	izinRequest.MessageTTD = "0"

	persetujuanResponse := pc.PimpinanService.PermohonanSetApproved(r.Context(), izinRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   persetujuanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Index implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexSPPD(w http.ResponseWriter, r *http.Request) {
	create := r.URL.Query().Get("c")

	result := strings.Compare(create, "true")

	if result == 0 {
		allUserResponse := pc.PimpinanService.GetAllUserId(r.Context())
		data := map[string]interface{}{
			"user": allUserResponse,
			"menu": "newPenugasan",
		}

		temp, err := template.ParseFiles("./view/pimpinan.html")
		helper.PanicIfError(err)

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	} else {
		allSuratResponse := pc.PimpinanService.SPPDGetAllSuratTugasJOINApprovedUser(r.Context())
		data := map[string]interface{}{
			"menu":   "sppd",
			"sppds":  allSuratResponse,
			"status": r.URL.Query().Get("v"),
		}

		temp, err := template.ParseFiles("./view/pimpinan.html")
		helper.PanicIfError(err)

		temp.Funcs(template.FuncMap{"index": helper.AddIndex})

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	}
}

// SPPDDetailSurat implements PimpinanController.
func (pc *PimpinanControllerImpl) SPPDDetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := pc.PimpinanService.SPPDGetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"surat": response,
		"menu":  "sppdView",
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

// IndexLaporan implements PimpinanController.
func (pc *PimpinanControllerImpl) IndexLaporan(w http.ResponseWriter, r *http.Request) {
	sppds := pc.PimpinanService.LaporanGetAllSPPD(r.Context())

	data := map[string]interface{}{
		"sppds": sppds,
		"menu":  "laporan",
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// LaporanDetail implements PimpinanController.
func (pc *PimpinanControllerImpl) LaporanDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["suratId"]
	idInt, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	surat := pc.PimpinanService.LaporanSPPDById(r.Context(), idInt)
	var data map[string]interface{}

	if surat.LaporanDokName != "" {
		data = map[string]interface{}{
			"sppd": surat,
			"menu": "viewLap",
			"lenP": len(surat.Participans),
		}
	}

	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// LaporanSetAprroved implements PimpinanController.
func (pc *PimpinanControllerImpl) LaporanSetAprroved(w http.ResponseWriter, r *http.Request) {
	approvedReq := laporanreqres.ApprovedLaporanRequest{
		UserId: 2,
	}

	helper.ReadFromRequestBody(r, &approvedReq)

	approvedRes := pc.PimpinanService.SetApprovedLaporan(r.Context(), approvedReq)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   approvedRes,
	}

	helper.WriteToResponseBody(w, response)
}

// Profile implements PimpinanController.
func (pc *PimpinanControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	userResponse := pc.PimpinanService.Profile(r.Context(), id)

	data := map[string]interface{}{
		"user": userResponse,
		"menu": "profile",
	}
	temp, err := template.ParseFiles("view/pimpinan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
