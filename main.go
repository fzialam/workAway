package main

import (
	"net/http"

	"github.com/fzialam/workAway/app"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	r := app.NewRouter(db, validate)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
