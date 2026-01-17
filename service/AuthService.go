package service

import (
	"fmt"
	"module/portofolio1/model"
	"module/portofolio1/repository"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Create(user model.SignIn) (model.SignIn, error)
	Login(data model.Login) (model.Login, error)
}

type Auth struct {
	repo     repository.Authrepository
	validate *validator.Validate
}

func NewAuthService(repo repository.Authrepository) AuthService {
	fmt.Println("Membuat instance service")
	return &Auth{
		repo:     repo,
		validate: validator.New(),
	}
}

func (as *Auth) Create(user model.SignIn) (model.SignIn, error) {
	fmt.Println("user servic create")
	fmt.Println(user)
	err := as.validate.Struct(user)
	if err != nil {
		fmt.Println("Error validating user:", err)
		return model.SignIn{}, err
	}

	fmt.Println("selesai proses validation")
	return user, nil
}

func (as *Auth) Login(data model.Login) (model.Login, error) {
	fmt.Println(data)

	user := model.Login{
		Email:    data.Email,
		Password: data.Password,
	}

	err := as.validate.Struct(user)
	if err != nil {
		fmt.Println("Error validating user:", err)
	}
	return data, nil
}
