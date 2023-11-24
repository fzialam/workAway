package app

import (
	"database/sql"
	"net/http"

	"github.com/fzialam/workAway/exception"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *mux.Router {
	r := mux.NewRouter()
	user := InitializedUser(db, validate)
	presensi := InitializedPresensi(db, validate)
	permohonan := InitializedPermohonan(db, validate)
	pimpinan := InitializedPimpinan(db, validate)
	// sppd := InitializedTU(db, validate)

	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./view/static/"))))

	r.HandleFunc("/login", user.Login).Methods("POST")
	r.HandleFunc("/login", user.IndexL).Methods("GET")
	r.HandleFunc("/register", user.Register).Methods("POST")
	r.HandleFunc("/register", user.IndexR).Methods("GET")
	r.HandleFunc("/logout", user.Logout).Methods("GET")

	// ===========> Pegawai Section Start <===========

	// Permohonan Section
	p := r.PathPrefix("/wp").Subrouter()
	p.HandleFunc("/permohonan/{userId}", permohonan.Index).Methods("GET")
	p.HandleFunc("/permohonan/{userId}", permohonan.CreatePermohonan).Methods("POST")

	// Mobile Section
	p.HandleFunc("/{userId}/mobile", presensi.GetSuratForPresensi).Methods("GET")
	p.HandleFunc("/{userId}/mobile", presensi.Presensi).Methods("POST")

	// ===========> Pegawai Section End <===========

	// ===========> TU Section Start <===========

	// Permohonan Section
	// tu := r.PathPrefix("/wt").Subrouter()
	// tu.HandleFunc("/sppd", sppd.Index).Methods("GET")
	// tu.HandleFunc("/sppd", sppd.CreateSPPD).Methods("POST")

	// ===========> TU Section End <===========

	// ===========> Pimpinan Section Start <===========

	// Persetujuan Section
	pp := r.PathPrefix("/wpp").Subrouter()
	pp.HandleFunc("/persetujuan", pimpinan.IndexPermohonan).Methods("GET")
	pp.HandleFunc("/{suratId}/persetujuan", pimpinan.PermohonanDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/persetujuan", pimpinan.PermohonanSetApproved).Methods("POST")

	// Penugasan Section
	pp.HandleFunc("/penugasan", pimpinan.IndexSPPD).Methods("GET")
	pp.HandleFunc("/penugasan", pimpinan.CreatePenugasan).Methods("POST")
	pp.HandleFunc("/{suratId}/penugasan", pimpinan.SPPDDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/penugasan", pimpinan.SPPDSetApproved).Methods("POST")

	// ===========> Pimpinan Section End <===========

	r.Use(exception.PanicHandler)

	return r
}
