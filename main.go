package main

import (
	"net/http"

	"github.com/fzialam/workAway/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := app.NewRouter()

	http.ListenAndServe(":3000", r)
}
