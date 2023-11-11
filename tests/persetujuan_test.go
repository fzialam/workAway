package tests

import (
	"context"
	"fmt"
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
	ctx := context.Background()
	sql, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	r, err := persetujuanrepository.NewPersetujuanRepo().GetAllSuratTugasJOINApprovedUser(ctx, sql)
	fmt.Println(r)
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
