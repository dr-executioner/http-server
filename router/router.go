package router

import (
	"http-server/handler"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/foobar", handler.FooHandler)
	return mux
}
