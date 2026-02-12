package controller

import (
	"encoding/json"
	"fmt"
	"module/portofolio1/errors"
	"module/portofolio1/model"
	"module/portofolio1/resources"
	"module/portofolio1/service"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type Auth struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *Auth {
	fmt.Println("Membuat instance controller")
	return &Auth{service: service}
}

func (ac *Auth) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login controller")

	email := r.FormValue("email")
	password := r.FormValue("password")

	service, err := ac.service.Login(model.Login{Email: email, Password: password})

	fmt.Println("err from service",err)

	fmt.Println("Service Login:", service)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(service)

}


// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func (ac *Auth) SignIn(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("auth-controller")

	ctx, span := tracer.Start(r.Context(), "SignIn-Controller" )
	defer span.End()

	start := time.Now()

	fmt.Println("sign in controller")

	var req model.SignIn

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.BadRequest(w, "Invalid request payload", nil)
		return
	}

	r.Body.Close()

	span.SetAttributes(
		attribute.String("email", req.Email),
	)

	fmt.Println("Email:", req.Email)
	fmt.Println("username:", req.Username)
	fmt.Println("password:", req.Password)

	user := model.SignIn{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	service, err := ac.service.Create(ctx, user)

	duration := time.Since(start)

	span.SetAttributes(
		attribute.Float64("duration_ms", float64(duration)),
	)

	fmt.Println(err)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		errors.BadRequest(w, "Validation Error", err)
		return
	}

	fmt.Println("data service dari controller: ",service)
	resources.Success(w, "User registered successfully", user)
}
