package presensicontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PresensiController interface {
	Presensi(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetSuratForPresensi(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
