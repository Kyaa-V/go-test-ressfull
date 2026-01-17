package router

import (
	"fmt"
	"module/portofolio1/controller"

	"github.com/go-chi/chi/v5"
)

func LoadAuthRouter(r chi.Router, authController *controller.Auth) {
	fmt.Println("Mendaftarkan route /login dan /signin")

	r.Post("/login", authController.Login)
	r.Post("/signin", authController.SignIn)
}
