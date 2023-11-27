package tests

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
)

func TestGetAllUserId(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/permohonan", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}
func TestCreatePermohonan(t *testing.T) {

	db := app.NewDB()
	router := setupRouter(db)
	// permohonanRequest := permohonanreqres.PermohonanRequest{
	// 	UserPemohonId:    1,
	// 	LokasiTujuan:     "UI",
	// 	JenisProgram:     "Mancing",
	// 	DokPendukungName: "header.Filename",
	// 	DokPendukungPdf:  "string(data)",
	// 	ParticipansId:    []int{1, 4, 6},
	// 	TglAwal:          time.Now().Add(24 * time.Hour),
	// 	TglAkhir:         time.Now().Add(96 * time.Hour),
	// }
	requestBody := strings.NewReader("lokasi=UI&jenis=Mancing&awal=2023-11-01&akhir=2023-12-10&participan=1&participan=4&participan=6")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/permohonan/1", requestBody)
	request.Header.Add("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryABC")

	x := []int{1, 4, 6}
	strSlice := make([]string, len(x))

	for i, v := range x {
		strSlice[i] = fmt.Sprint(v)
	}

	result := "["
	result += strings.Join(strSlice, ", ")
	result += "]"

	request.Form = url.Values{}

	request.Form.Add("lokasi", "UI")
	request.Form.Add("jenis", "Mancing")
	request.Form.Add("participans", result)
	request.Form.Add("TglAwal", "2023-11-02")
	request.Form.Add("TglAkhir", "2023-12-10")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)
	fmt.Println(string(bd))
}

func TestAddSurat(t *testing.T) {
	tx, err := app.NewDB().Begin()
	ctx := context.Background()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	x, err := pegawairepository.NewPegawaiRepo().CreateSurat(ctx, tx, entity.SuratTugas{
		LokasiTujuan:     "Unesa",
		JenisProgram:     "1",
		DokPendukungName: "dokumne.pdf",
		DokPendukungPdf:  "ISI DOKUMEN PDF",
		TglAwal:          "2023-10-12",
		TglAkhir:         "2023-12-12",
	})
	fmt.Println(x)
	row := tx.QueryRowContext(ctx, "SELECT MAX(id) AS `last_index` FROM `surat_tugas`;")
	helper.PanicIfError(err)
	var su entity.SuratTugas
	err = row.Scan(
		&su.Id,
		&su.LokasiTujuan,
		&su.JenisProgram,
		&su.DokPendukungName,
		&su.DokPendukungPdf,
		&su.TglAwal,
		&su.TglAkhir,
		&su.CreateAt)
	helper.PanicIfError(err)
	log.Println(su)

}
