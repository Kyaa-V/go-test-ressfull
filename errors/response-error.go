package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  bool 	  `json:"status"`
	Message string   `json:"message"`
	ErrorCode string  `json:"error_code"`
	Details any      `json:"details,omitempty"`
}

func Error(w http.ResponseWriter, statusCode int, message string, errorCode string, details any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := ErrorResponse{
		Status: false,
		Message: message,
		ErrorCode: errorCode,
		Details: details,
	}

	json.NewEncoder(w).Encode(response)
}

func BadRequest(w http.ResponseWriter, message string, details any) {
    Error(w, http.StatusBadRequest, message, "ErrBadRequest", details)
}

func Unauthorized(w http.ResponseWriter, message string) {
    Error(w, http.StatusUnauthorized, message, "ErrUnauthorized", nil)
}

func Forbidden(w http.ResponseWriter, message string) {
    Error(w, http.StatusForbidden, message, "ErrForbidden", nil)
}

func NotFound(w http.ResponseWriter, message string) {
    Error(w, http.StatusNotFound, message, "ErrNotFound", nil)
}

func ValidationError(w http.ResponseWriter, message string, details any) {
    Error(w, http.StatusBadRequest, message, "ErrValidation", details)
}

func InternalServerError(w http.ResponseWriter, message string) {
    Error(w, http.StatusInternalServerError, message, "ErrInternalServer", nil)
}