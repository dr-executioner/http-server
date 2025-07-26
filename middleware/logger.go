package middleware

import (
	"log"
	"net/http"
	"os"
)

func SetupLogger(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not open file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)
	log.Println("=====Server Started======")
}
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})

}
