package main

import (
	"log"
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

	log.Println("Server start ...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
	log.Println("Server running at port :3000")
}
