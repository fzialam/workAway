package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fzialam/workAway/app"
)

func TestPresensiSucces(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	truncateTsble(db, "presensi")
	// img := ImageToBase64()

	reqBody := strings.NewReader(`{
		"user_id" : 1,
		"surat_tugas_id" : 2,
		"gambar" : "img",
		"lokasi": "LOKASI"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/w/mobile/1", reqBody)
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
	truncateTsble(db, "presensi")
	// img := ImageToBase64()

	// Surat Tugas ditolak
	reqBody := strings.NewReader(`{
		"user_id" : 1,
		"surat_tugas_id" : 2,
		"gambar" : "img",
		"lokasi": ""
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/wp/1/mobile", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}

func TestGetSuratSucces(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/mobile/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}

func TestGetSuratFailed(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/mobile/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}
