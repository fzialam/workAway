package presensicontroller

import (
	"net/http"
)

type PresensiController interface {
	Presensi(w http.ResponseWriter, r *http.Request)
	GetSuratForPresensi(w http.ResponseWriter, r *http.Request)
}
