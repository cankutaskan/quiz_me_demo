package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ParseQueryParam(param string, defaultValue int) int {
	if param == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(param)
	if err != nil || parsed <= 0 {
		return defaultValue
	}
	return parsed
}

func EncodeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
