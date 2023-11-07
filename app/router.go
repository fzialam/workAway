package app

import (
	"database/sql"

	"github.com/fzialam/workAway/exception"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	r := httprouter.New()
	user := InitializedUser(db, validate)
	presensi := InitializedPresensi(db, validate)
	permohonan := InitializedPermohonan(db, validate)

	// r.GET("/", controller.Index)
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.GET("/mobile/:userId", presensi.GetSuratForPresensi)
	r.POST("/mobile/:userId", presensi.Presensi)
	r.GET("/permohonan/:userId", permohonan.Index)
	r.POST("/permohonan/:userId", permohonan.CreatePermohonan)
	r.GET("/all-user", user.FindAll)

	r.PanicHandler = exception.ErrorHandler

	return r
}
