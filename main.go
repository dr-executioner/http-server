package main

import (
	"http-server/middleware"
	"http-server/router"
	"log"
	"net/http"
)

func main() {

	middleware.SetupLogger("server.log")
	mux := router.SetupRouter()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)
	loggerHandler := middleware.Logger(mux)
	handlerWithCORS := middleware.AllowCors(loggerHandler)

	addr := ":6969"
	log.Printf("Server running at %s\n", addr)
	err := http.ListenAndServe(addr, handlerWithCORS)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
