package app

import (
	"github.com/fzialam/workAway/controller"
	presensicontroller "github.com/fzialam/workAway/controller/presensi_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	"github.com/fzialam/workAway/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(user usercontroller.UserController, presensi presensicontroller.PresensiController) *httprouter.Router {
	r := httprouter.New()

	// Define an endpoint to fetch all data
	r.GET("/", controller.Index)
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)
	r.POST("/mobile", presensi.Presensi)
	r.GET("/mobile", presensi.GetSurat)
	r.GET("/all-user", user.FindAll)

	r.PanicHandler = exception.ErrorHandler

	return r
}
