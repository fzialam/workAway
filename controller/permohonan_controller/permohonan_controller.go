package permohonancontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PermohonanController interface {
	Index(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CreatePermohonan(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
