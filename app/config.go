package app

import (
	"database/sql"

	permohonancontroller "github.com/fzialam/workAway/controller/permohonan_controller"
	pimpinancontroller "github.com/fzialam/workAway/controller/pimpinan_controller"
	presensicontroller "github.com/fzialam/workAway/controller/presensi_controller"
	tucontroller "github.com/fzialam/workAway/controller/tu_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	permohonanrepository "github.com/fzialam/workAway/repository/permohonan_repository"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
	presensirepository "github.com/fzialam/workAway/repository/presensi_repository"
	turepository "github.com/fzialam/workAway/repository/tu_repository"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	permohonanservice "github.com/fzialam/workAway/service/permohonan_service"
	pimpinanservice "github.com/fzialam/workAway/service/pimpinan_service"
	presensiservice "github.com/fzialam/workAway/service/presensi_service"
	tuservice "github.com/fzialam/workAway/service/tu_service"
	userservice "github.com/fzialam/workAway/service/user_service"
	"github.com/go-playground/validator/v10"
)

func InitializedPresensi(db *sql.DB, validate *validator.Validate) presensicontroller.PresensiController {
	presensiRepo := presensirepository.NewPresensiRepo()
	presensiService := presensiservice.NewPresensiService(presensiRepo, db, validate)
	presensiController := presensicontroller.NewPresensiController(presensiService)
	return presensiController
}

func InitializedUser(db *sql.DB, validate *validator.Validate) usercontroller.UserController {
	userRepo := userrepository.NewUserRepo()
	userService := userservice.NewUserService(userRepo, db, validate)
	userController := usercontroller.NewUserController(userService)

	return userController
}

func InitializedPermohonan(db *sql.DB, validate *validator.Validate) permohonancontroller.PermohonanController {
	permohonanRepo := permohonanrepository.NewPermohonanRepo()
	permohonanService := permohonanservice.NewPermohonanService(permohonanRepo, db, validate)
	permohonanController := permohonancontroller.NewPermohonanController(permohonanService)

	return permohonanController
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
