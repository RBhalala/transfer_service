package handler

import (
	"net/http"
	"encoding/json"
	"strings"
)

// To modify the error msg and pass it as JSON response
func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	// Setting fallback error msg
	if strings.TrimSpace(message) == "" {
		message = "Action Invalid"
	}

	resp := map[string]string{"error": message}
	json.NewEncoder(w).Encode(resp)
}