package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
	pegawaiservice "github.com/fzialam/workAway/service/pegawai_service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db := app.NewDB()
	ctx := context.Background()

	// tx, err := db.Begin()
	// helper.PanicIfError(err)

	// defer helper.CommitOrRollback(tx)

	pr := pegawairepository.NewPegawaiRepo()
	ps := pegawaiservice.NewPegawaiService(pr, db, validator.New())

	hasil := ps.LaporanGetAllSPPDByUserId(ctx, 1)

	for _, x := range hasil {
		fmt.Printf("x.Message: %v\n", x.Message)
		fmt.Printf("x.Message: %v\n", len(x.Message))
	}
}
