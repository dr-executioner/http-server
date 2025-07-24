package handler

import (
	"http-server/utils"
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
	utils.WriteJSON(w, http.StatusOK, response)
}
