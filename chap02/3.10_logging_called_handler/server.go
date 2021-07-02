package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello!\n")
}

// 呼び出されたハンドラをロギングする
// http.HandlerFunc インタフェースの定義はコレ
// type HandlerFunc func(ResponseWriter, *Request)
func log(h http.HandlerFunc) http.HandlerFunc {
	// 無名関数を返す
	return func(w http.ResponseWriter, r *http.Request) {
		// https://pkg.go.dev/runtime#FuncForPC
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)

		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", log(hello))
	_ = server.ListenAndServe()
}
