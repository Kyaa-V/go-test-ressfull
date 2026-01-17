package application

import (
	"fmt"
	"net/http"

	"module/portofolio1/controller"
)

type App struct {
	router http.Handler
}

func New(authController *controller.Auth) *App {
	return &App{
		router: setupRouter(authController),
	}
}

func (a *App) Start() error {
	fmt.Println("Server running on :3000")
	return http.ListenAndServe(":3000", a.router)
}
