package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
)

func TestIndex(t *testing.T) {

	db := app.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pr, err := pimpinanrepository.NewPimpinanRepo().IndexPenugasan(ctx, tx)
	helper.PanicIfError(err)

	for _, st := range pr {
		fmt.Printf("st.Id: %v\n", st.Id)
		fmt.Printf("st.Rincian.Id: %v\n", st.Rincian.Id)
	}

}

func TestNullFullAnggaran(t *testing.T) {
	db := app.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	err = pimpinanrepository.NewPimpinanRepo().SetNullFullAnggaran(ctx, tx, 10)
	helper.PanicIfError(err)
	fmt.Printf("err: %v\n", err)
}
