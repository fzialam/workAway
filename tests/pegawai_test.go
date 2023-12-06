package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
	pegawaiservice "github.com/fzialam/workAway/service/pegawai_service"
	"github.com/go-playground/validator/v10"
)

func TestGetlaporanByUserId(t *testing.T) {
	db := app.NewDB()

	v := validator.New()
	pr := pegawairepository.NewPegawaiRepo()

	ctx := context.Background()
	surats := pegawaiservice.NewPegawaiService(pr, db, v).LaporanGetAllSPPDByUserId(ctx, 1)

	for _, stjlr := range surats {
		fmt.Printf("stjlr.StatusPimpinan: %v\n", stjlr.StatusPimpinan)
		fmt.Printf("stjlr.StatusKeuangan: %v\n", stjlr.StatusKeuangan)
	}
}

func TestIndexPegawai(t *testing.T) {
	db := app.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pr, err := pegawairepository.NewPegawaiRepo().Index(ctx, tx, 1)
	helper.PanicIfError(err)

	fmt.Printf("pr: %v\n", pr)

}
