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
	pegawai := InitializedPegawai(db, validate)
	pimpinan := InitializedPimpinan(db, validate)
	tu := InitializedTU(db, validate)

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
	p.HandleFunc("/{userId}/permohonan", pegawai.Index).Methods("GET")
	p.HandleFunc("/{userId}/permohonan", pegawai.CreatePermohonan).Methods("POST")

	// p.HandleFunc("/{userId}/penugasan", pegawai.PenugasanIndex).Methods("GET")

	p.HandleFunc("/{userId}/laporan", pegawai.LaporanIndex).Methods("GET")
	p.HandleFunc("/{userId}/laporan-ak", pegawai.UploadLapAktivitas).Methods("POST")
	p.HandleFunc("/{userId}/laporan-ang", pegawai.UploadLapAnggaran).Methods("POST")

	// Mobile Section
	p.HandleFunc("/{userId}/mobile", pegawai.GetSurat).Methods("GET")
	p.HandleFunc("/{userId}/mobile", pegawai.Presensi).Methods("POST")

	// ===========> Pegawai Section End <===========

	// ===========> TU Section Start <===========

	// Permohonan Section
	t := r.PathPrefix("/wt").Subrouter()
	t.HandleFunc("/sppd", tu.Index).Methods("GET")
	t.HandleFunc("/{suratId}/sppd", tu.DetailSurat).Methods("GET")
	t.HandleFunc("/{suratId}/sppd", tu.CreateSPPD).Methods("POST")

	// ===========> TU Section End <===========

	// ===========> Pimpinan Section Start <===========

	// Persetujuan Section
	pp := r.PathPrefix("/wpp").Subrouter()
	pp.HandleFunc("/permohonan", pimpinan.IndexPermohonan).Methods("GET")
	pp.HandleFunc("/{suratId}/permohonan", pimpinan.PermohonanDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/permohonan", pimpinan.PermohonanSetApproved).Methods("POST")

	// Penugasan Section
	pp.HandleFunc("/sppd", pimpinan.IndexSPPD).Methods("GET")
	pp.HandleFunc("/penugasan", pimpinan.CreatePenugasan).Methods("POST")
	pp.HandleFunc("/{suratId}/sppd", pimpinan.SPPDDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/sppd", pimpinan.SPPDSetApproved).Methods("POST")

	pp.HandleFunc("/laporan", pimpinan.IndexLaporan).Methods("GET")
	pp.HandleFunc("/{suratId}/laporan", pimpinan.LaporanDetail).Methods("GET")
	pp.HandleFunc("/{suratId}/laporan", pimpinan.LaporanSetAprroved).Methods("POST")
	// ===========> Pimpinan Section End <===========

	r.Use(exception.PanicHandler)

	return r
}
