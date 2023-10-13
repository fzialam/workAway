package tests

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/fzialam/workAway/helper"
	"github.com/fzialam/workAway/model/entity"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "rootsql:@tcp(localhost:3306)/workaway")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := entity.User{
		Email:    "email@unesa.ac.id",
		Password: "password",
	}
	SQL := "SELECT `id`, `nip`, `rank`, `email`, `password` FROM `user` WHERE email=? AND password=?"
	fmt.Println(user)
	rows, err := db.Query(SQL, user.Email, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	newUser := entity.User{}
	if rows.Next() {
		err := rows.Scan(&newUser.Id, &newUser.NIP, &newUser.Rank, &newUser.Email, &newUser.Password)
		helper.PanicIfError(err)
		fmt.Println(newUser, nil)
	} else {
		fmt.Println(newUser, errors.New("user is not found"))
	}
}
