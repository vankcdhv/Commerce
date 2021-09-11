package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Api helper initialized")
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}