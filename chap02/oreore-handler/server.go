package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	myHandler := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myHandler,
	}

	server.ListenAndServe()
}
