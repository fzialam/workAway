package tests

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("mysql", "rootsql:@tcp(localhost:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Sukses DB")
}
