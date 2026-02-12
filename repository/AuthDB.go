package repository

import (
	"context"
	"database/sql"
	"fmt"
	"module/portofolio1/model"

	_ "github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/otel"
)

type Authrepository interface {
	Create(ctx context.Context,user model.SignIn) (model.SignIn, error)
	FindByEmail(user model.Login) (model.Login, error)
	Login(data model.Login) (model.Login, error)
}

type AuthDB struct {
	DB *sql.DB
}

func NewAuthrepository(db *sql.DB) *AuthDB {
	fmt.Println("Membuat instance AuthDB repository")
	return &AuthDB{DB: db}
}

func (a *AuthDB) Login(data model.Login) (model.Login, error) {
	query := "SELECT email, password FROM users WHERE email = ? AND password = ?"
	row := a.DB.QueryRow(query, data.Email, data.Password)

	var result model.Login
	err := row.Scan(&result.Email, &result.Password)
	if err != nil {
		return model.Login{}, err
	}

	return result, nil
}

func (a *AuthDB) Create(ctx context.Context, user model.SignIn) (model.SignIn, error) {
	tracer := otel.Tracer("auth-repository")
	meter := otel.Meter("repo-meter")

	requestCounter, _ := meter.Int64Counter("https_request_total")

	_, span := tracer.Start(ctx, "repository(Create)")
	defer span.End()

	query := "INSERT INTO users (username,email, password) VALUES (?, ?, ?)"
	_, err := a.DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		span.RecordError(err)
		return model.SignIn{}, err
	}
	requestCounter.Add(ctx, 1, )
	return user, nil
}

func (a *AuthDB) FindByEmail(user model.Login) (model.Login, error) {
	query := "SELECT email, password FROM users WHERE email = ?"
	row := a.DB.QueryRow(query, user.Email)

	var result model.Login
	err := row.Scan(&result.Email, &result.Password)
	if err != nil {
		return model.Login{}, err
	}
	return result, nil
}
