package app

import (
	"database/sql"

	keuangancontroller "github.com/fzialam/workAway/controller/keuangan_controller"
	pegawaicontroller "github.com/fzialam/workAway/controller/pegawai_controller"
	pimpinancontroller "github.com/fzialam/workAway/controller/pimpinan_controller"
	tucontroller "github.com/fzialam/workAway/controller/tu_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	keuanganrepository "github.com/fzialam/workAway/repository/keuangan_repository"
	pegawairepository "github.com/fzialam/workAway/repository/pegawai_repository"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
	turepository "github.com/fzialam/workAway/repository/tu_repository"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	keuanganservice "github.com/fzialam/workAway/service/keuangan_service"
	pegawaiservice "github.com/fzialam/workAway/service/pegawai_service"
	pimpinanservice "github.com/fzialam/workAway/service/pimpinan_service"
	tuservice "github.com/fzialam/workAway/service/tu_service"
	userservice "github.com/fzialam/workAway/service/user_service"
	"github.com/go-playground/validator/v10"
)

func InitializedUser(db *sql.DB, validate *validator.Validate) usercontroller.UserController {
	userRepo := userrepository.NewUserRepo()
	userService := userservice.NewUserService(userRepo, db, validate)
	userController := usercontroller.NewUserController(userService)

	return userController
}

func InitializedPimpinan(db *sql.DB, validate *validator.Validate) pimpinancontroller.PimpinanController {
	pr := pimpinanrepository.NewPimpinanRepo()
	ps := pimpinanservice.NewPimpinanService(pr, db, validate)
	pc := pimpinancontroller.NewPimpinanController(ps)

	return pc
}

func InitializedTU(db *sql.DB, validate *validator.Validate) tucontroller.TUController {
	tr := turepository.NewTURepo()
	ts := tuservice.NewTUService(tr, db, validate)
	tc := tucontroller.NewTUController(ts)

	return tc
}

func InitializedPegawai(db *sql.DB, validate *validator.Validate) pegawaicontroller.PegawaiController {
	pr := pegawairepository.NewPegawaiRepo()
	ps := pegawaiservice.NewPegawaiService(pr, db, validate)
	pc := pegawaicontroller.NewPegawaiController(ps)

	return pc
}

func InitializedKeuangan(db *sql.DB, validate *validator.Validate) keuangancontroller.KeuanganController {
	kr := keuanganrepository.NewKeuanganRepo()
	ks := keuanganservice.NewKeuanganService(kr, db, *validate)
	kc := keuangancontroller.NewKeuanganController(ks)

	return kc
}
