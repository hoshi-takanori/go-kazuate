package main

import (
	"net/http"
	"os"
)

func main() {
	server := NewServer()
	go server.Start()

	http.Handle("/ws", server.WebSocketHandler())
	http.Handle("/", http.FileServer(http.Dir("public")))

	// cert.pem and key.pem are generated by
	// go run /usr/local/go/src/crypto/tls/generate_cert.go --host hostname
	_, err1 := os.Stat("cert.pem")
	_, err2 := os.Stat("key.pem")
	if err1 == nil && err2 == nil {
		http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}
