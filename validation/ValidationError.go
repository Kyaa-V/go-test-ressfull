package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ConvertValdationError(err error) map[string]string {

	errorMsg := map[string]string{}
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			fmt.Println("Field:", e.Field(), "Tag:", e.Tag())
			switch e.Field() {
				case "Email":
					errorMsg[e.Field()] = "Email tidak valid atau sudah terdaftar"
				case "Username":
					errorMsg[e.Field()] = "Username tidak valid atau sudah terdaftar"
				case "Password":
					errorMsg[e.Field()] = "Password minimal 6-20 karakter"
				default:
					errorMsg[e.Field()] = "Data yang anda masukkan salah"
			}
		}
	}
	return errorMsg
}