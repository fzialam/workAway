package pegawaicontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
	laporanreqres "github.com/fzialam/workAway/model/req_res/laporan_req_res"
	permohonanreqres "github.com/fzialam/workAway/model/req_res/permohonan_req_res"
	presensireqres "github.com/fzialam/workAway/model/req_res/presensi_req_res"
	surattugasreqres "github.com/fzialam/workAway/model/req_res/surat_tugas_req_res"
	pegawaiservice "github.com/fzialam/workAway/service/pegawai_service"
	"github.com/gorilla/mux"
)

type PegawaiControllerImpl struct {
	PegawaiService pegawaiservice.PegawaiService
}

func NewPegawaiController(pegawaiService pegawaiservice.PegawaiService) PegawaiController {
	return &PegawaiControllerImpl{
		PegawaiService: pegawaiService,
	}
}

// Index implements PegawaiController.
func (pc *PegawaiControllerImpl) Index(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idS := vars["userId"]
	id, _ := strconv.Atoi(idS)

	index, err := pc.PegawaiService.Index(r.Context(), id)
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"menu":   "home",
		"userId": id,
		"index":  index,
	}
	temp, err := template.ParseFiles("./view/pegawai.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// Index implements PegawaiController.
func (pc *PegawaiControllerImpl) IndexPermohonan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idS := vars["userId"]
	id, _ := strconv.Atoi(idS)

	var data map[string]interface{}
	create := r.URL.Query().Get("c")
	view := r.URL.Query().Get("v")

	if create == "true" {
		allUserResponse := pc.PegawaiService.GetAllUserId(r.Context(), id)
		data = map[string]interface{}{
			"menu":   "newPermohonan",
			"userId": id,
			"data":   allUserResponse,
		}
		temp, err := template.ParseFiles("./view/pegawai.html")
		helper.PanicIfError(err)

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	} else if view != "" {
		idSurat, err := strconv.Atoi(view)
		helper.PanicIfError(err)

		respons := pc.PegawaiService.GetSuratById(r.Context(), idSurat)

		data = map[string]interface{}{
			"menu":       "viewPermohonan",
			"userId":     id,
			"permohonan": respons,
			"lenP":       len(respons.Participans),
		}
		temp, err := template.ParseFiles("./view/pegawai.html")
		helper.PanicIfError(err)

		temp.Funcs(template.FuncMap{"index": helper.AddIndex})

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	} else {

		request := surattugasreqres.GetSuratRequest{
			UserId: id,
			Tipe:   "permohonan",
		}

		permohonan := pc.PegawaiService.GetSurat(r.Context(), request)

		suratIds := ""

		for i, v := range permohonan {
			if i+1 == len(permohonan) {
				suratIds += strconv.Itoa(v.Id)
			} else {
				suratIds += strconv.Itoa(v.Id) + ","
			}
		}

		data = map[string]interface{}{
			"menu":       "permohonan",
			"permohonan": permohonan,
			"userId":     id,
			"suratIds":   suratIds,
		}
		temp, err := template.ParseFiles("./view/pegawai.html")
		helper.PanicIfError(err)

		temp.Funcs(template.FuncMap{"index": helper.AddIndex})

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	}
}

// CreatePermohonan implements PegawaiController.
func (pc *PegawaiControllerImpl) CreatePermohonan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	permohonanRequest := permohonanreqres.PermohonanRequest{
		UserPemohonId: userId,
	}
	helper.ReadFromRequestBody(r, &permohonanRequest)

	permohonanResponse := pc.PegawaiService.CreatePermohonan(r.Context(), permohonanRequest)

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   permohonanResponse,
	}
	helper.WriteToResponseBody(w, response)
}

// Presensi implements PegawaiController.
func (pc *PegawaiControllerImpl) Presensi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	set := r.URL.Query().Get("set")

	presensiRequest := presensireqres.PresensiFotoRequest{
		UserId: userId,
	}

	helper.ReadFromRequestBody(r, &presensiRequest)
	helper.PanicIfError(err)

	if set == "" {
		presensiResponse := pc.PegawaiService.PresensiFoto(r.Context(), presensiRequest)
		response := model.Response{
			Code:   200,
			Status: "OK",
			Data:   presensiResponse,
		}

		helper.WriteToResponseBody(w, response)
	} else {
		presensiRequest.Set = "set"
		presensiResponse := pc.PegawaiService.PresensiFoto(r.Context(), presensiRequest)
		response := model.Response{
			Code:   200,
			Status: "OK",
			Data:   presensiResponse,
		}

		helper.WriteToResponseBody(w, response)

	}
}

// GetSuratPresensi implements PegawaiController.
func (pc *PegawaiControllerImpl) GetSuratPresensi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	getSuratResponse := pc.PegawaiService.GetSuratPresensi(r.Context(), userId)

	helper.WriteToResponseBody(w, getSuratResponse)
}

// IndexSPPD implements PegawaiController.
func (pc *PegawaiControllerImpl) IndexSPPD(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	getSuratReq := surattugasreqres.GetSuratRequest{
		UserId: userId,
		Tipe:   "sppd",
	}
	sppds := pc.PegawaiService.GetSurat(r.Context(), getSuratReq)

	data := map[string]interface{}{
		"menu":   "sppd",
		"sppds":  sppds,
		"userId": userId,
	}
	temp, err := template.ParseFiles("./view/pegawai.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// IndexLaporan implements PegawaiController.
func (pc *PegawaiControllerImpl) IndexLaporan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	var data map[string]interface{}

	SPPDIdS := r.URL.Query().Get("id")

	if SPPDIdS != "" {
		sppdId, err := strconv.Atoi(SPPDIdS)
		helper.PanicIfError(err)

		sppd := pc.PegawaiService.LaporanGetSPPDById(r.Context(), laporanreqres.LaporanGetSPPDByIdRequest{
			UserId:       userId,
			SuratTugasId: sppdId,
		})

		data = map[string]interface{}{
			"menu":   "uploadLap",
			"sppd":   sppd,
			"id":     SPPDIdS,
			"userId": userId,
			"lenP":   len(sppd.Participans),
		}
	} else if SPPDIdS == "" {
		surats := pc.PegawaiService.LaporanGetAllSPPDByUserId(r.Context(), userId)

		suratIds := ""

		for i, v := range surats {
			if i+1 == len(surats) {
				suratIds += strconv.Itoa(v.Id)
			} else {
				suratIds += strconv.Itoa(v.Id) + ","
			}
		}

		data = map[string]interface{}{
			"menu":     "laporan",
			"surats":   surats,
			"id":       SPPDIdS,
			"userId":   userId,
			"suratIds": suratIds,
		}

	}
	temp, err := template.ParseFiles("./view/pegawai.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// UploadLapAktivitas implements PegawaiController.
func (pc *PegawaiControllerImpl) UploadLapAktivitas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	suratIdS := r.URL.Query().Get("id")
	suratId, err := strconv.Atoi(suratIdS)
	helper.PanicIfError(err)

	laporanReq := laporanreqres.UploadLaporanRequest{
		UserId:       userId,
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &laporanReq)

	laporanResponse := pc.PegawaiService.UploadLapAktivitas(r.Context(), laporanReq)

	response := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   laporanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// UploadLapAnggaran implements PegawaiController.
func (pc *PegawaiControllerImpl) UploadLapAnggaran(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	suratIdS := r.URL.Query().Get("id")
	suratId, err := strconv.Atoi(suratIdS)
	helper.PanicIfError(err)

	laporanReq := laporanreqres.UploadLaporanRequest{
		UserId:       userId,
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &laporanReq)

	laporanResponse := pc.PegawaiService.UploadLapAnggaran(r.Context(), laporanReq)

	response := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   laporanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// SetLapAktivitas implements PegawaiController.
func (pc *PegawaiControllerImpl) SetLapAktivitas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	suratIdS := r.URL.Query().Get("id")
	suratId, err := strconv.Atoi(suratIdS)
	helper.PanicIfError(err)

	laporanReq := laporanreqres.UploadLaporanRequest{
		UserId:       userId,
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &laporanReq)

	laporanResponse := pc.PegawaiService.SetLapAktivitas(r.Context(), laporanReq)

	response := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   laporanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// SetLapAnggaran implements PegawaiController.
func (pc *PegawaiControllerImpl) SetLapAnggaran(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	suratIdS := r.URL.Query().Get("id")
	suratId, err := strconv.Atoi(suratIdS)
	helper.PanicIfError(err)

	laporanReq := laporanreqres.UploadLaporanRequest{
		UserId:       userId,
		SuratTugasId: suratId,
	}

	helper.ReadFromRequestBody(r, &laporanReq)

	laporanResponse := pc.PegawaiService.SetLapAnggaran(r.Context(), laporanReq)

	response := model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   laporanResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// Profile implements PegawaiController.
func (pc *PegawaiControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	userResponse := pc.PegawaiService.Profile(r.Context(), userId)

	data := map[string]interface{}{
		"user":   userResponse,
		"menu":   "profile",
		"userId": userId,
	}
	temp, err := template.ParseFiles("./view/pegawai.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
