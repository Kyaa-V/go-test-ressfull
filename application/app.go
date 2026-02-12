package application

import (
	"context"
	"fmt"
	"module/portofolio1/controller"
	"module/portofolio1/middleware"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":3000",
		Handler: middleware.TelemetriMiddleware(a.router),
		BaseContext: func(net.Listener) context.Context {return ctx},
	}

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	fmt.Println("Server running on :3000")
	select {
		case err := <-srvErr:
			return err
		case <-ctx.Done():
			fmt.Println("Shutting down server...")
			stop()
	}

	err := srv.Shutdown(context.Background())

	return err
}
