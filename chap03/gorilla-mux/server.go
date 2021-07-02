package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	router := mux.NewRouter()

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	router.HandleFunc("/hello", hello)
	server.ListenAndServe()
}
