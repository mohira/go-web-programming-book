package main

import (
	"fmt"
	"net/http"
)

// HelloHandler は http.Handler インターフェイスを満たしている
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello!")
}

// ハンドラの実行をロギングする
func log(h http.Handler) http.Handler {
	// http.HandlerFunc型で包む
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

// ハンドラを実行する前にユーザーの権限を確認する
func protect(h http.Handler) http.Handler {
	// http.HandlerFunc型で包む
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ここで認証かなにかのコードを実行するとする
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}

	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))

	_ = server.ListenAndServe()
}
