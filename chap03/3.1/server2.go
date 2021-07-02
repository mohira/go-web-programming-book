package main

import "net/http"

// http.ListenAndServe() じゃなくて http.Server を直接使う場合
func main() {
	server := http.Server{Addr: "", Handler: nil}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
