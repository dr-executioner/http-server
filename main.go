package main

import (
	"http-server/router"
	"log"
	"net/http"
)

func main() {

	mux := router.SetupRouter()

	addr := ":6969"
	log.Printf("Server running at %s\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
