package main

import (
	"module/portofolio1/application"
	"module/portofolio1/controller"
	"module/portofolio1/repository"
	"module/portofolio1/service"
)

func main() {
	db := application.NewConnection()
	if db != nil {
		panic("Tidak dapat koneksi ke database")
	}

	authRepo := repository.NewAuthrepository(application.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	app := application.New(authController)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
