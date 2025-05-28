package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func teachersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(""))
	case http.MethodPost:
		w.Write([]byte(""))
	case http.MethodPut:
		w.Write([]byte(""))
	case http.MethodPatch:
		w.Write([]byte(""))
	case http.MethodDelete:
		w.Write([]byte(""))
	}
	w.Write([]byte("teachers root"))
}

func studentsHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("students root"))
}

func main() {

	port := ":9000"

	cert := "../../cert.pem"
	key := "../../key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Route root"))
	})

	mux.HandleFunc("/teachers", teachersHandlers)

	mux.HandleFunc("/students", studentsHandlers)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server running on port: ", port)
	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatal("Error starting the server", err)
	}

}
