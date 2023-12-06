package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
)

func TestIndex(t *testing.T) {

	db := app.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pr := pimpinanrepository.NewPimpinanRepo().GetAllFotoParticipanById(ctx, tx, entity.ParticipanJoinUser{
		UserId:       6,
		SuratTugasId: 33,
	})
	helper.PanicIfError(err)

	fmt.Printf("pr: %v\n", pr)

}
