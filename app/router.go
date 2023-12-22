package app

import (
	"database/sql"
	"net/http"

	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *mux.Router {
	r := mux.NewRouter()
	user := InitializedUser(db, validate)
	pegawai := InitializedPegawai(db, validate)
	pimpinan := InitializedPimpinan(db, validate)
	tu := InitializedTU(db, validate)
	keuangan := InitializedKeuangan(db, validate)
	admin := InitializedAdmin(db, validate)

	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./view/static/"))))

	r.HandleFunc("/login", user.IndexL).Methods("GET")
	r.HandleFunc("/login", user.Login).Methods("POST")
	r.HandleFunc("/register", user.Register).Methods("POST")
	r.HandleFunc("/register", user.IndexR).Methods("GET")
	r.HandleFunc("/logout", user.Logout).Methods("GET")

	// ===========> Pegawai Section Start <===========

	// Permohonan Section
	p := r.PathPrefix("/wp").Subrouter()
	p.HandleFunc("/{userId}/home", pegawai.Index).Methods("GET")
	p.HandleFunc("/{userId}/permohonan", pegawai.IndexPermohonan).Methods("GET")
	p.HandleFunc("/{userId}/permohonan", pegawai.CreatePermohonan).Methods("POST")

	p.HandleFunc("/{userId}/sppd", pegawai.IndexSPPD).Methods("GET")

	p.HandleFunc("/{userId}/laporan", pegawai.IndexLaporan).Methods("GET")
	p.HandleFunc("/{userId}/laporan-ak", pegawai.UploadLapAktivitas).Methods("POST")
	p.HandleFunc("/{userId}/laporan-ang", pegawai.UploadLapAnggaran).Methods("POST")

	p.HandleFunc("/{userId}/set-laporan-ak", pegawai.SetLapAktivitas).Methods("POST")
	p.HandleFunc("/{userId}/set-laporan-ang", pegawai.SetLapAnggaran).Methods("POST")

	p.HandleFunc("/{userId}/profile", pegawai.Profile).Methods("GET")
	p.HandleFunc("/{userId}/profile", user.UpdateProfile).Methods("POST")
	p.HandleFunc("/{userId}/password", user.ChangePassword).Methods("POST")
	p.HandleFunc("/{userId}/image", user.ChangeImage).Methods("POST")

	// Mobile Section
	r.HandleFunc("/{userId}/mobile", pegawai.GetSuratPresensi).Methods("GET")
	r.HandleFunc("/{userId}/mobile", pegawai.Presensi).Methods("POST")

	p.Use(middleware.VerifyRoleUser(0))

	// ===========> Pegawai Section End <===========

	// ===========> TU Section Start <===========

	// Permohonan Section
	t := r.PathPrefix("/wt").Subrouter()
	t.HandleFunc("/home", tu.Index).Methods("GET")
	t.HandleFunc("/sppd", tu.IndexSPPD).Methods("GET")
	t.HandleFunc("/{suratId}/sppd", tu.DetailSurat).Methods("GET")
	t.HandleFunc("/{suratId}/sppd", tu.CreateSPPD).Methods("POST")

	t.HandleFunc("/profile", tu.Profile).Methods("GET")
	t.HandleFunc("/profile", user.UpdateProfile).Methods("POST")
	t.HandleFunc("/password", user.ChangePassword).Methods("POST")
	t.HandleFunc("/image", user.ChangeImage).Methods("POST")

	t.Use(middleware.VerifyRoleUser(2))

	// ===========> TU Section End <===========

	// ===========> Pimpinan Section Start <===========

	// Persetujuan Section
	pp := r.PathPrefix("/wpp").Subrouter()
	pp.HandleFunc("/home", pimpinan.Index).Methods("GET")
	pp.HandleFunc("/permohonan", pimpinan.IndexPermohonan).Methods("GET")
	pp.HandleFunc("/{suratId}/permohonan", pimpinan.PermohonanDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/permohonan", pimpinan.PermohonanSetApproved).Methods("POST")

	// Penugasan Section
	pp.HandleFunc("/sppd", pimpinan.IndexSPPD).Methods("GET")
	pp.HandleFunc("/penugasan", pimpinan.IndexPenugasan).Methods("GET")
	pp.HandleFunc("/penugasan", pimpinan.CreatePenugasan).Methods("POST")
	pp.HandleFunc("/sppd", pimpinan.IndexSPPD).Methods("GET")
	pp.HandleFunc("/{suratId}/sppd", pimpinan.SPPDDetailSurat).Methods("GET")
	pp.HandleFunc("/{suratId}/sppd", pimpinan.SPPDSetApproved).Methods("POST")

	pp.HandleFunc("/laporan", pimpinan.IndexLaporan).Methods("GET")
	pp.HandleFunc("/{suratId}/laporan", pimpinan.LaporanDetail).Methods("GET")
	pp.HandleFunc("/{suratId}/laporan", pimpinan.LaporanSetAprroved).Methods("POST")

	pp.HandleFunc("/profile", pimpinan.Profile).Methods("GET")
	pp.HandleFunc("/profile", user.UpdateProfile).Methods("POST")
	pp.HandleFunc("/password", user.ChangePassword).Methods("POST")
	pp.HandleFunc("/image", user.ChangeImage).Methods("POST")

	pp.Use(middleware.VerifyRoleUser(1))
	// ===========> Pimpinan Section End <===========

	// ===========> Bagian Keuangan Section Start <===========
	bk := r.PathPrefix("/wk").Subrouter()
	bk.HandleFunc("/home", keuangan.Index).Methods("GET")
	bk.HandleFunc("/rincian-biaya", keuangan.IndexPermohonan).Methods("GET")
	bk.HandleFunc("/{suratId}/rincian-biaya", keuangan.IndexPermohonan).Methods("GET")
	bk.HandleFunc("/{suratId}/rincian-biaya", keuangan.UploadRincian).Methods("POST")
	bk.HandleFunc("/{suratId}/set-rincian", keuangan.SetRincian).Methods("POST")

	bk.HandleFunc("/sppd", keuangan.IndexSPPD).Methods("GET")
	bk.HandleFunc("/sppd", keuangan.SetFullAnggaran).Methods("POST")

	bk.HandleFunc("/laporan", keuangan.IndexLaporan).Methods("GET")
	bk.HandleFunc("/laporan", keuangan.SetApprovedLaporan).Methods("POST")

	bk.HandleFunc("/profile", keuangan.Profile).Methods("GET")
	bk.HandleFunc("/profile", user.UpdateProfile).Methods("POST")
	bk.HandleFunc("/password", user.ChangePassword).Methods("POST")
	bk.HandleFunc("/image", user.ChangeImage).Methods("POST")

	bk.Use(middleware.VerifyRoleUser(3))

	wa := r.PathPrefix("/wa").Subrouter()
	wa.HandleFunc("/home", admin.Index).Methods("GET")
	wa.HandleFunc("/permohonan", admin.Permohonan).Methods("GET")
	wa.HandleFunc("/penugasan", admin.Penugasan).Methods("GET")
	wa.HandleFunc("/laporan-ak", admin.LapAKK).Methods("GET")
	wa.HandleFunc("/laporan-ag", admin.LapAGG).Methods("GET")
	wa.HandleFunc("/user", admin.UserGET).Methods("GET")
	wa.HandleFunc("/{userId}/user", admin.UserGETById).Methods("GET")
	wa.HandleFunc("/{userId}/user", admin.UserPUT).Methods("POST")

	wa.HandleFunc("/profile", admin.Profile).Methods("GET")
	wa.HandleFunc("/profile", user.UpdateProfile).Methods("POST")
	wa.HandleFunc("/password", user.ChangePassword).Methods("POST")
	wa.HandleFunc("/image", user.ChangeImage).Methods("POST")

	wa.Use(middleware.VerifyRoleUser(4))

	// ===========> Bagian Keuangan Section End <===========

	r.Use(exception.PanicHandler)

	return r
}
