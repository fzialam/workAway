package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/fzialam/workAway/helper"
	adminrepository "github.com/fzialam/workAway/repository/admin_repository"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db := app.NewDB()
	ctx := context.Background()

	tx, err := db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	pr, err := adminrepository.NewAdminRepo().UserGET(ctx, tx)
	helper.PanicIfError(err)

	// fmt.Printf("pr: %v\n", pr)
	for i, stj := range pr {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("stj.Id: %v\n", stj.Id)
	}
}
