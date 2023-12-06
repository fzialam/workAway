package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	turepository "github.com/fzialam/workAway/repository/tu_repository"
)

func TestDetailSurat(t *testing.T) {
	db := app.NewDB()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	stjaup, err := turepository.NewTURepo().GetSuratTugasById(context.Background(), tx, 32)
	helper.PanicIfError(err)

	fmt.Printf("stjaup.Id: %v\n", stjaup.Id)
	fmt.Printf("stjaup.Status: %v\n", stjaup.Status)
	fmt.Printf("stjaup.OtherStatus: %v\n", stjaup.OtherStatus)

}
