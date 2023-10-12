package main

import (
	"log"

	"github.com/fzialam/workAway/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := app.NewRouter()

	// Start the API server
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
