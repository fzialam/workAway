package tests

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	persetujuanrepository "github.com/fzialam/workAway/repository/persetujuan_repository"
)

func TestGetSuratWithJOINById(t *testing.T) {
	ctx := context.Background()
	sql, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	r, err := persetujuanrepository.NewPersetujuanRepo().GetSuratTugasById(ctx, sql, 1)
	fmt.Println(r)
}
func TestGetParticippanWithJOINBySuratId(t *testing.T) {
	ctx := context.Background()
	sql, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	r, err := persetujuanrepository.NewPersetujuanRepo().GetAllParticipanJOINUserBySuratId(ctx, sql, 31)
	fmt.Println(r)
}

func TestGetAllSuratTugasJOINApprovedUser(t *testing.T) {
	db := app.NewDB()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/w/persetujuan", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bd, _ := io.ReadAll(response.Body)

	fmt.Println(string(bd))
}

func TestSetApproved(t *testing.T) {
	ctx := context.Background()
	sql, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(sql)
	w := time.Now().Format("2006-01-02 15:04:05")
	x := entity.Izin{
		SuratTugasId: 1,
		Status:       2,
		CreateAt:     w,
	}
	r := persetujuanrepository.NewPersetujuanRepo().SetApproved(ctx, sql, x)
	fmt.Println(r)
}

func TestTime(t *testing.T) {
	log.Println(helper.TimeNowToString())
}
