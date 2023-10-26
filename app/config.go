package app

import (
	"database/sql"

	presensicontroller "github.com/fzialam/workAway/controller/presensi_controller"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	presensirepository "github.com/fzialam/workAway/repository/presensi_repository"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
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
