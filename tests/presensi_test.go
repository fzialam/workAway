package tests

import (
	"log"
	"net/http"
	"testing"

	"github.com/fzialam/workAway/app"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func init() {
	db := app.NewDB()
	validate := validator.New()
	presInit := app.InitializedPresensi(db, validate)
	r := httprouter.New()
	r.POST("/mobile", presInit.Presensi)

	log.Println("Server start ...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
	log.Println("Server running at port :3000")
}

func TestPresensiSucces(t *testing.T) {

}
func TestPresensiFailed(t *testing.T) {

}
