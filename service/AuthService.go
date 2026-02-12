package service

import (
	"context"
	"fmt"
	"module/portofolio1/model"
	"module/portofolio1/repository"
	"module/portofolio1/validation"

	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type ValidationError struct {
	Errors map[string]string
}

func (ve *ValidationError) Error() string {
	return "validation error"
}

type AuthService interface {
	Create(ctx context.Context, user model.SignIn) (model.SignIn, error)
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

func (as *Auth) Create(ctx context.Context, user model.SignIn) (model.SignIn, error) {
	tracer :=  otel.Tracer("auth-service")

	ctx, span := tracer.Start(ctx, "service(Create)")
	defer span.End()

	fmt.Println("user servic create")
	err := as.validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		customErr := validation.ConvertValdationError(err)
		fmt.Println("customErr:", customErr)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return model.SignIn{}, &ValidationError{Errors: customErr}
	}

	fmt.Println("selesai proses validation")
	fmt.Println("start reepo create")

	user, err = as.repo.Create(ctx, user)

	fmt.Println("selesai proses repo create")

	if err != nil {
		return model.SignIn{}, err
	}

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
