package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "World!")
}
