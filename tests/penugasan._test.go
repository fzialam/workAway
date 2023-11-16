package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	penugasanrepository "github.com/fzialam/workAway/repository/penugasan_repository"
)

func TestGetAllSuratPenugasan(t *testing.T) {

	tx, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	response, err := penugasanrepository.NewPenugasanRepo().GetAllSuratTugasJOINApprovedUser(
		context.Background(),
		tx,
	)
	helper.PanicIfError(err)
	for _, res := range response {
		fmt.Println(res.Id)
	}
}

func TestGetSuratTugasById(t *testing.T) {
	tx, err := app.NewDB().Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	response, err := penugasanrepository.NewPenugasanRepo().GetSuratTugasById(
		context.Background(),
		tx,
		1,
	)
	helper.PanicIfError(err)
	fmt.Println(response.DokumenName)
}
