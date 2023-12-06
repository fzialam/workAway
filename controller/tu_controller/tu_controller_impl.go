package tucontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model"
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
	index := tc.TUService.IndexTU(r.Context())
	temp, err := template.ParseFiles("view/tu.html")
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"menu":  "home",
		"index": index,
	}
	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	temp.Execute(w, data)
}

// IndexSPPD implements TUController.
func (tc *TUControllerImpl) IndexSPPD(w http.ResponseWriter, r *http.Request) {
	surats := tc.TUService.GetAllSuratTugasJOINApprovedUser(r.Context())
	temp, err := template.ParseFiles("view/tu.html")
	helper.PanicIfError(err)

	data := map[string]interface{}{
		"menu":   "sppd",
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

	response := model.Response{
		Code:   200,
		Status: "OK",
		Data:   cs,
	}

	helper.WriteToResponseBody(w, response)
}

// DetailSurat implements TUController.
func (tc *TUControllerImpl) DetailSurat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idSurat, err := strconv.Atoi(vars["suratId"])
	helper.PanicIfError(err)

	response := tc.TUService.GetSuratTugasById(r.Context(), idSurat)

	data := map[string]interface{}{
		"surat": response,
		"lenP":  len(response.Participans),
		"menu":  "sppdView",
	}

	temp, err := template.ParseFiles("view/tu.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}

// Profile implements TUController.
func (tc *TUControllerImpl) Profile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("userId")
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Value)
	helper.PanicIfError(err)

	userResponse := tc.TUService.Profile(r.Context(), id)

	data := map[string]interface{}{
		"user": userResponse,
		"menu": "profile",
	}

	temp, err := template.ParseFiles("./view/tu.html")
	helper.PanicIfError(err)

	temp.Funcs(template.FuncMap{"index": helper.AddIndex})

	err = temp.Execute(w, data)
	helper.PanicIfError(err)
}
