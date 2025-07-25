package handler

import (
	"encoding/json"
	"fmt"
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
		utils.WriterError(w, http.StatusMethodNotAllowed, "Method Not ALlowed")
		return
	}

	var req EchoRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Message == "" {
		utils.WriterError(w, http.StatusBadRequest, "Invalid JSON or misson body")
		return
	}
	fmt.Println(r.Body)
	resp := EchoResponse{Received: req.Message}
	utils.WriteJSON(w, http.StatusOK, resp)
}
