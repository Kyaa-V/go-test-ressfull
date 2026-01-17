package application

import (
	"net/http"

	"module/portofolio1/controller"
	"module/portofolio1/router"

	"github.com/go-chi/chi/v5"
)

func setupRouter(authController *controller.Auth) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Route("/api", func(r chi.Router) {
		router.LoadAuthRouter(r, authController)
	})

	return r
}
