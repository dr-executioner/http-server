package handler

import (
	"encoding/json"
	"http-server/utils"
	"net/http"
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Received string `json:"received"`
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriterError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var req EchoRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Message == "" {
		utils.WriterError(w, http.StatusBadRequest, "Invalid JSON or missing body")
		return
	}
	resp := EchoResponse{Received: req.Message}
	utils.WriteJSON(w, http.StatusOK, resp)
}
