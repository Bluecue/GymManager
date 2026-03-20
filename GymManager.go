package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	http.Handle("/", http.FileServer(http.Dir("./resources/")))
	port := ":8443"

	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(port, "auth/cert.pem", "auth/key.pem", nil)
	log.Fatal(err)

}
