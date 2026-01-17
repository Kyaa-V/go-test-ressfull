package resources

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, message string, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	// fmt.Println(json.NewEncoder(w).Encode(response))
	json.NewEncoder(w).Encode(response)
}