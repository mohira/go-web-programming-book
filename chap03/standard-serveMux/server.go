package main

import (
	"fmt"
	"net/http"
)

func main() {
	myMux := http.ServeMux{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myMux,
	}

	myMux.HandleFunc("/hello", hello)
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
