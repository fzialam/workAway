package app

import (
	"database/sql"

	penugasancontroller "github.com/fzialam/workAway/controller/penugasan_controller"
	permohonancontroller "github.com/fzialam/workAway/controller/permohonan_controller"
	persetujuancontroller "github.com/fzialam/workAway/controller/persetujuan_controller"
	pimpinancontroller "github.com/fzialam/workAway/controller/pimpinan_controller"
	presensicontroller "github.com/fzialam/workAway/controller/presensi_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	penugasanrepository "github.com/fzialam/workAway/repository/penugasan_repository"
	permohonanrepository "github.com/fzialam/workAway/repository/permohonan_repository"
	persetujuanrepository "github.com/fzialam/workAway/repository/persetujuan_repository"
	pimpinanrepository "github.com/fzialam/workAway/repository/pimpinan_repository"
	presensirepository "github.com/fzialam/workAway/repository/presensi_repository"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	penugasanservice "github.com/fzialam/workAway/service/penugasan_service"
	permohonanservice "github.com/fzialam/workAway/service/permohonan_service"
	persetujuanservice "github.com/fzialam/workAway/service/persetujuan_service"
	pimpinanservice "github.com/fzialam/workAway/service/pimpinan_service"
	presensiservice "github.com/fzialam/workAway/service/presensi_service"
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

func InitializedPersetujuan(db *sql.DB, validate *validator.Validate) persetujuancontroller.PersetujunanController {
	pr := persetujuanrepository.NewPersetujuanRepo()
	ps := persetujuanservice.NewPersetujuanService(pr, db, validate)
	pc := persetujuancontroller.NewPersetujuanController(ps)

	return pc
}
func InitializedPenugasan(db *sql.DB, validate *validator.Validate) penugasancontroller.PenugasanController {
	pr := penugasanrepository.NewPenugasanRepo()
	ps := penugasanservice.NewPenugasanService(pr, db, validate)
	pc := penugasancontroller.NewPenugasanController(ps)

	return pc
}

func InitializedPimpinan(db *sql.DB, validate *validator.Validate) pimpinancontroller.PimpinanController {
	pr := pimpinanrepository.NewPimpinanRepo()
	ps := pimpinanservice.NewPimpinanService(pr, db, validate)
	pc := pimpinancontroller.NewPimpinanController(ps)

	return pc
}
