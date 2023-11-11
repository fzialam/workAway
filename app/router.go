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
	persetujuan := InitializedPersetujuan(db, validate)

	staticDir := "/static/"
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./view/static/"))))

	r.HandleFunc("/login", user.Login).Methods("POST")
	r.HandleFunc("/login", user.IndexL).Methods("GET")
	r.HandleFunc("/register", user.Register).Methods("POST")
	r.HandleFunc("/register", user.IndexR).Methods("GET")
	r.HandleFunc("/logout", user.Logout).Methods("GET")

	s := r.PathPrefix("/w").Subrouter()
	// s.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./view/static/"))))
	// Mobile Section
	s.HandleFunc("/mobile/{userId}", presensi.GetSuratForPresensi).Methods("GET")
	s.HandleFunc("/mobile/{userId}", presensi.Presensi).Methods("POST")

	// Permohonan Section
	s.HandleFunc("/permohonan/{userId}", permohonan.Index).Methods("GET")
	s.HandleFunc("/permohonan/{userId}", permohonan.CreatePermohonan).Methods("POST")

	// Penugasan Section
	s.HandleFunc("/penugasan/{userId}", permohonan.Index)
	s.HandleFunc("/penugasan/{userId}", permohonan.CreatePermohonan).Methods("POST")

	// Persetujuan Section
	s.HandleFunc("/persetujuan", persetujuan.Index).Methods("GET")

	s.HandleFunc("/all-user", user.FindAll).Methods("GET")

	r.Use(exception.PanicHandler)
	// s.Use()

	return r
}
