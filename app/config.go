package app

import (
	"database/sql"

	permohonancontroller "github.com/fzialam/workAway/controller/permohonan_controller"
	presensicontroller "github.com/fzialam/workAway/controller/presensi_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	permohonanrepository "github.com/fzialam/workAway/repository/permohonan_repository"
	presensirepository "github.com/fzialam/workAway/repository/presensi_repository"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	permohonanservice "github.com/fzialam/workAway/service/permohonan_service"
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
