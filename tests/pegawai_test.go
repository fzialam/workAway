package tests

import (
	"context"
	"log"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/exception"
	"github.com/fzialam/workAway/helper"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
)

func TestGetlaporanByUserId(t *testing.T) {
	tx, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()
	surat, err := pegawairepository.NewPegawaiRepo().LaporanGetAllSPPDByUserId(ctx, tx, 1)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	log.Println(surat[len(surat)-1].TglAwal)
	log.Println(surat[len(surat)-1].CreateAt)
}
