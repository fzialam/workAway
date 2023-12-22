package keuangancontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	approvedreqres "github.com/fzialam/workAway/model/req_res/approved_req_res"
	keuanganreqres "github.com/fzialam/workAway/model/req_res/keuangan_req_res"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	keuanganservice "github.com/fzialam/workAway/service/keuangan_service"
	"github.com/gorilla/mux"
)

type KeuanganControllerImpl struct {
	KeuanganService keuanganservice.KeuanganService
}

func NewKeuanganController(keuanganService keuanganservice.KeuanganService) KeuanganController {
	return &KeuanganControllerImpl{
		KeuanganService: keuanganService,
	}
}

// Index implements KeuanganController.
func (kc *KeuanganControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	index, err := kc.KeuanganService.Index(r.Context())
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"index": index,
		"menu":  "home",
	}

	temp, err := template.ParseFiles("./view/keuangan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// IndexKeuangan implements KeuanganController.
func (kc *KeuanganControllerImpl) IndexPermohonan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["suratId"]
	view := r.URL.Query().Get("v")

	var data map[string]interface{}

	if (idS == "") && (view == "") {
		surats := kc.KeuanganService.ListPermohonanApproved(r.Context())
		data = map[string]interface{}{
			"surats": surats,
			"menu":   "permohonan",
		}
	} else if (idS != "") && (view != "") {
		suratId, _ := strconv.Atoi(idS)
		surat := kc.KeuanganService.PermohonanApprovedById(r.Context(), suratId)
		data = map[string]interface{}{
			"surat": surat,
			"menu":  "permohonanDetail",
			"lenP":  len(surat.Participans),
		}
	}
	temp, err := template.ParseFiles("./view/keuangan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// UploadRincian implements KeuanganController.
func (kc *KeuanganControllerImpl) UploadRincian(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["suratId"]
	suratId, _ := strconv.Atoi(idS)
	uploadRequest := keuanganreqres.UploadRincianAnggaranRequest{
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &uploadRequest)
	uploadResponse := kc.KeuanganService.UploadRincianBiaya(r.Context(), uploadRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   uploadResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// UploadRincian implements KeuanganController.
func (kc *KeuanganControllerImpl) SetRincian(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["suratId"]
	suratId, _ := strconv.Atoi(idS)
	uploadRequest := keuanganreqres.UploadRincianAnggaranRequest{
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &uploadRequest)
	uploadResponse := kc.KeuanganService.SetRincian(r.Context(), uploadRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   uploadResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// IndexSPPD implements KeuanganController.
func (kc *KeuanganControllerImpl) IndexSPPD(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	surats := kc.KeuanganService.ListSPPDApproved(r.Context())
	data = map[string]interface{}{
		"surats": surats,
		"menu":   "sppd",
	}

	temp, err := template.ParseFiles("./view/keuangan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// SetFullAnggaran implements KeuanganController.
func (kc *KeuanganControllerImpl) SetFullAnggaran(w http.ResponseWriter, r *http.Request) {
	approvedReq := approvedreqres.ApprovedRequest{}
	helper.ReadFromRequestBody(r, &approvedReq)

	approvedResponse := kc.KeuanganService.SetFullAnggaran(r.Context(), approvedReq)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   approvedResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// IndexLaporan implements KeuanganController.
func (kc *KeuanganControllerImpl) IndexLaporan(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	surats := kc.KeuanganService.ListLaporanSPPD(r.Context())
	data = map[string]interface{}{
		"surats": surats,
		"menu":   "laporan",
	}

	temp, err := template.ParseFiles("./view/keuangan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// SetApprovedLaporan implements KeuanganController.
func (kc *KeuanganControllerImpl) SetApprovedLaporan(w http.ResponseWriter, r *http.Request) {
	approvedLaporanReq := laporanreqres.ApprovedLaporanRequest{
		UserId: 4,
	}
	helper.ReadFromRequestBody(r, &approvedLaporanReq)

	approvedResponse := kc.KeuanganService.SetApprovedLaporan(r.Context(), approvedLaporanReq)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   approvedResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Profile implements KeuanganController.
func (kc *KeuanganControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	userResponse := kc.KeuanganService.Profile(r.Context(), id)

	data := map[string]interface{}{
		"user": userResponse,
		"menu": "profile",
	}

	temp, err := template.ParseFiles("./view/keuangan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
