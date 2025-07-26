package handler

import (
	"fmt"
	"http-server/utils"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriterError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	name := r.URL.Query().Get("name")
	fmt.Printf("%#v\n", r)
	if name == "" {
		name = "Guest"
	}

	msg := fmt.Sprintf("Hello, %s!", name)
	response := HelloResponse{Message: msg}
	utils.WriteJSON(w, http.StatusOK, response)

}
