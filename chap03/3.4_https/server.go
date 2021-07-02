package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// cert.pem: SSL証明書
	// key.pem: 秘密鍵
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		panic(err)
	}
}
