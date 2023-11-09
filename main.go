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
	// userInit := app.InitializedUser(db, validate)
	// presensiinit := app.InitializedPresensi(db, validate)
	r := app.NewRouter(db, validate)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("/view/static")))

	// Mengarahkan semua permintaan yang dimulai dengan "/static/" ke file server
	r.PathPrefix("/static/").Handler(staticHandler)

	log.Println("Server running at port http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
