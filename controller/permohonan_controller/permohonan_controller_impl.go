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
	/*r.ParseMultipartForm(5 * 1024 * 1024) // Batasan ukuran file 5 MB
	// Mendapatkan file yang diunggah
	file, header, err := r.FormFile("pendukung")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// Baca konten file PDF
	pdfBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Gagal membaca file PDF", http.StatusInternalServerError)
		return
	}

	// Konversi ke base64
	pdfBase64 := base64.StdEncoding.EncodeToString(pdfBytes)

	// Mendapatkan data dari elemen select (Participan)
	var participans []int
	participanValues := r.Form["participan"]
	for _, value := range participanValues {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		participans = append(participans, intValue)
	}

	userIdPemohon, _ := strconv.Atoi(p.ByName("userId"))
	jenis, _ := strconv.Atoi(p.ByName("jenis"))
	permohonanRequest := permohonanreqres.PermohonanRequest{
		UserPemohonId:    userIdPemohon,
		LokasiTujuan:     r.Form.Get("lokasi"),
		JenisProgram:     jenis,
		DokPendukungName: header.Filename,
		DokPendukungPdf:  pdfBase64,
		ParticipansId:    participans,
		TglAwal:          r.FormValue("awal"),
		TglAkhir:         r.FormValue("akhir"),
	}

	// permohonanResponse := pc.PermohonanService.CreatePermohonan(r.Context(), permohonanRequest)
	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   permohonanRequest,
	}
	fmt.Fprintln(w, participanValues)
	helper.WriteToResponseBody(w, response)*/

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
