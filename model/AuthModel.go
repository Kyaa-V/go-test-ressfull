package model

type SignIn struct {
	Username string `json:"username" validate:"required,min=3,max=20"` 
	Email 	string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}