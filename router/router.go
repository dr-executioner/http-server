package router

import (
	"http-server/handler"
	"http-server/utils"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/info", handler.InfoHandler)
	mux.HandleFunc("/foobar", handler.FooHandler)
	mux.HandleFunc("/echo", handler.EchoHandler)
	mux.HandleFunc("/notes/", handler.GetNoteById)
	mux.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetNotesHandler(w, r)
		case http.MethodPost:
			handler.CreateNoteHandler(w, r)
		default:
			utils.WriterError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}

	})

	return mux
}
