package repository

import (
	"database/sql"
	"fmt"
	"module/portofolio1/model"

	_ "github.com/go-sql-driver/mysql"
)

type Authrepository interface {
	Create(user model.SignIn) (model.SignIn, error)
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

func (a *AuthDB) Create(user model.SignIn) (model.SignIn, error) {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err := a.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		return model.SignIn{}, err
	}
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
