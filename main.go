package main

import (
	"net/http"

	"github.com/fzialam/workAway/app"
	usercontroller "github.com/fzialam/workAway/controller/user_controller"
	userrepository "github.com/fzialam/workAway/repository/user_repository"
	userservice "github.com/fzialam/workAway/service/user_service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepo := userrepository.NewUserRepo()
	userService := userservice.NewUserService(userRepo, db, validate)
	userController := usercontroller.NewUserController(userService)
	r := app.NewRouter(userController)

	http.ListenAndServe(":3000", r)
}
