package controller

import (
	"encoding/json"
	"fmt"
	"module/portofolio1/error"
	"module/portofolio1/model"
	"module/portofolio1/resources"
	"module/portofolio1/service"
	"net/http"
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

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Service Login:", service)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(service)

}

func (ac *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign in controller")

	var req model.SignIn

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		error.BadRequest(w, "Invalid request payload", nil)
		return
	}

	r.Body.Close()

	fmt.Println("Email:", req.Email)
	fmt.Println("username:", req.Username)
	fmt.Println("password:", req.Password)

	user := model.SignIn{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	service, err := ac.service.Create(user)

	fmt.Println("data service dari controller: ",service)
	resources.Success(w, "User registered successfully", user)
}
