package main

import (
	"log"
	"net/http"

	"github.com/fzialam/workAway/app"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userInit := app.InitializedUser(db, validate)
	presensiinit := app.InitializedPresensi(db, validate)
	r := app.NewRouter(userInit, presensiinit)

	log.Println("Server start ...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
	log.Println("Server running at port :3000")
}
