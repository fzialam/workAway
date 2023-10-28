package tests

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	presInit := app.InitializedPresensi(db, validate)
	r := httprouter.New()
	r.POST("/mobile", presInit.Presensi)

	r.PanicHandler = exception.ErrorHandler
	return r
}

func truncateTsble(db *sql.DB) {
	tableName := "presensi"
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

func TestPresensiSucces(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	truncateTsble(db)
	// img := ImageToBase64()

	reqBody := strings.NewReader(`{
		"user_id" : 1,
		"surat_tugas_id" : 1,
		"gambar" : "img",
		"lokasi": "LOKASI"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/mobile", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))

}
func TestPresensiFailed(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	truncateTsble(db)
	// img := ImageToBase64()

	// Surat Tugas ditolak
	reqBody := strings.NewReader(`{
		"user_id" : 1,
		"surat_tugas_id" : 2,
		"gambar" : "img",
		"lokasi": "LOKASI"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/mobile", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}
