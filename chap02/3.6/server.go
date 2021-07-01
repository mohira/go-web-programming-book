package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// どのパスでアクセスしても同じ結果
// $ curl 127.0.0.1:8080
// Hello World!
// $ curl 127.0.0.1:8080/foo/bar
// Hello World!
// $ curl 127.0.0.1:8080/anything/at/all
// Hello World!
