package handler

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Status  string `json:"status"`
	Message string `json:message"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := InfoResponse{
		Status:  "OK",
		Message: "Server is running",
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}

}
