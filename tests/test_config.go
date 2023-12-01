package tests

import (
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ResponseData struct {
	Code   int                      `json:"code"`
	Status string                   `json:"status"`
	Data   []map[string]interface{} `json:"data"`
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	pegawai := app.InitializedPegawai(db, validate)
	tu := app.InitializedTU(db, validate)
	r := mux.NewRouter()
	s := r.PathPrefix("/w").Subrouter()
	s.HandleFunc("/mobile/{userId}", pegawai.GetSurat).Methods("GET")
	s.HandleFunc("/mobile/{userId}", pegawai.Presensi).Methods("POST")
	s.HandleFunc("/permohonan/{userId}", pegawai.CreatePermohonan).Methods("POST")
	s.HandleFunc("/pengajuan/{userId}", pegawai.IndexPermohonan).Methods("GET")

	r.HandleFunc("/sppd", tu.Index).Methods("GET")
	r.HandleFunc("/{suratId}/sppd", tu.DetailSurat).Methods("GET")
	r.HandleFunc("/{suratId}/sppd", tu.CreateSPPD).Methods("POST")

	r.Use(exception.PanicHandler)
	return r
}

func truncateTsble(db *sql.DB, tableName string) {
	sql := "TRUNCATE " + tableName
	_, err := db.Exec(sql)
	helper.PanicIfError(err)
}

func ImageToBase64() string {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile("./erd_workaway.png")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	encodBase64 := base64.StdEncoding.EncodeToString(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64," + encodBase64
	case "image/png":
		base64Encoding = "data:image/png;base64," + encodBase64
	}
	// Print the full base64 representation of the image
	return base64Encoding
}
