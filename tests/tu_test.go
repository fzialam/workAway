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

func TestDetailSurat(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{
		"dokumen_name" : "NAMA DOkumen",
		"dokumen_pdf" : "PDF DOKUMEN"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/32/sppd", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))

}
