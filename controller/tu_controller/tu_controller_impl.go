package tucontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	tureqres "github.com/fzialam/workAway/model/req_res/tu_req_res"
	tuservice "github.com/fzialam/workAway/service/tu_service"
	"github.com/gorilla/mux"
)

func NewTUController(tuService tuservice.TUService) TUController {
	return &TUControllerImpl{
		TUService: tuService,
	}
}

type TUControllerImpl struct {
	TUService tuservice.TUService
}

// Index implements TUController.
func (tc *TUControllerImpl) Index(w http.ResponseWriter, r *http.Request) {
	surats := tc.TUService.GetAllSuratTugasJOINApprovedUser(r.Context())
	temp, err := template.ParseFiles("view/tu.html")
	helper.PanicIfError(err)
	var data map[string]interface{}

	data = map[string]interface{}{
		"status": r.URL.Query().Get("v"),
		"surats": surats,
	}
	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	temp.Execute(w, data)
}

// CreateSPPD implements TUController.
func (tc *TUControllerImpl) CreateSPPD(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	suratId, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	createReq := tureqres.CreateSPPDRequest{
		SuratTugasId: suratId,
	}
	helper.ReadFromRequestBody(r, &createReq)

	cs := tc.TUService.CreateSPPD(r.Context(), createReq)

	helper.WriteToResponseBody(w, cs)
}

// DetailSurat implements TUController.
func (tc *TUControllerImpl) DetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := tc.TUService.GetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"surat":  response,
		"lenP":   len(response.Participans),
		"status": r.URL.Query().Get("v"),
	}

	temp, err := template.ParseFiles("view/tu.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
