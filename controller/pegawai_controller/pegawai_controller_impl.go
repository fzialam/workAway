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

	var data map[string]interface{}
	create := r.URL.Query().Get("c")
	status := r.URL.Query().Get("v")

	if (create == "true") && (status == "") {
		allUserResponse := pc.PegawaiService.GetAllUserId(r.Context())
		data = map[string]interface{}{
			"create": create,
			"data":   allUserResponse,
		}
	} else if (status == "lap") && (create == "") {
		data = map[string]interface{}{
			"status": status,
			"data":   nil,
		}
	}
	temp, err := template.ParseFiles("./view/pegawai.html")
	helper.PanicIfError(err)

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
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

	presensiRequest := presensireqres.PresensiFotoRequest{
		UserId: userId,
	}
	helper.ReadFromRequestBody(r, &presensiRequest)

	presensiResponse := pc.PegawaiService.PresensiFoto(r.Context(), presensiRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   presensiResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// GetSurat implements PegawaiController.
func (pc *PegawaiControllerImpl) GetSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)

	getSuratRequest := presensireqres.GetSuratForPresensiRequest{
		UserId: userId,
	}
	getSuratResponse := pc.PegawaiService.GetSurat(r.Context(), getSuratRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   getSuratResponse,
	}

	helper.WriteToResponseBody(w, response)
}

// LaporanIndex implements PegawaiController.
func (pc *PegawaiControllerImpl) LaporanIndex(w http.ResponseWriter, r *http.Request) {
	// panic("unimplemented")
	vars := mux.Vars(r)
	userIdS := vars["userId"]
	userId, err := strconv.Atoi(userIdS)
	helper.PanicIfError(err)
	var data map[string]interface{}

	SPPDIdS := r.URL.Query().Get("id")

	temp, err := template.ParseFiles("./view/pegawai-laporan.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	if SPPDIdS != "" {
		sppdId, err := strconv.Atoi(SPPDIdS)
		helper.PanicIfError(err)
		sppd := pc.PegawaiService.LaporanGetSPPDById(r.Context(), laporanreqres.LaporanGetSPPDByIdRequest{
			UserId:       userId,
			SuratTugasId: sppdId,
		})

		data = map[string]interface{}{
			"sppd":    sppd,
			"id":      SPPDIdS,
			"userIdS": userIdS,
			"lenP":    len(sppd.Participans),
		}

		err = temp.Execute(w, data)
		helper.PanicIfError(err)

	} else {
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
			"surats":   surats,
			"id":       SPPDIdS,
			"userIdS":  userIdS,
			"suratIds": suratIds,
		}

		err = temp.Execute(w, data)
		helper.PanicIfError(err)
	}
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
