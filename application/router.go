package application

import (
	"net/http"
	// "os"
	"module/portofolio1/controller"
	"module/portofolio1/middleware"
	"module/portofolio1/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/go-chi/chi/v5"
	_ "module/portofolio1/docs"
)

func setupRouter(authController *controller.Auth) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// r.Get("swagger/docs.json", func(w http.ResponseWriter, r *http.Request) {
	// 	jsonFile, _ := os.ReadFile("../docs/swagger.json")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(jsonFile)
	// })

	r.Get("swagger/*", httpSwagger.Handler(
		httpSwagger.URL("https://localhost:3000/swagger/doc.json"),
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Route("/api", func(r chi.Router) {
		router.LoadAuthRouter(r, authController)
	})

	return r
}
