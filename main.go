package main

import (
	"http-server/middleware"
	"http-server/router"
	"log"
	"net/http"
)

func main() {

	mux := router.SetupRouter()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	handlerWithCORS := middleware.AllowCors(mux)

	addr := ":6969"
	log.Printf("Server running at %s\n", addr)
	err := http.ListenAndServe(addr, handlerWithCORS)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
