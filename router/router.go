package router

import (
	"http-server/handler"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/info", handler.InfoHandler)
	mux.HandleFunc("/foobar", handler.FooHandler)
	return mux
}
