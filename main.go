package main

import (
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/api", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, HTTP/3 and WebTransport!"))
		},
	)

	// Handler for WebTransport
	mux.HandleFunc(
		"/webtransport", func(w http.ResponseWriter, r *http.Request) {
			// Handle WebTransport logic here
			w.Write([]byte("WebTransport endpoint"))
		},
	)

	server := http3.Server{
		Addr:    ":443",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServeTLS("/etc/ssl/certs/nginx-selfsigned.crt", "/etc/ssl/private/nginx-selfsigned.key"))
}
