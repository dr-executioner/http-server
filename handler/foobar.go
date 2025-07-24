package handler

import (
	"fmt"
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowded on this route", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "FooBar")
}
