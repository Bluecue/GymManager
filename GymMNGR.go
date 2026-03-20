package main

import (
	// "html/template"
	// "github.com/gin-gonic/gin"
	// "io"
	"log"
	"net/http"
)

//	func main() {
//		fs := http.FileServer(http.Dir("./resources/"))
//		http.Handle("/resources/", http.StripPrefix("/resources/", fs))
//
//		log.Println("Serving on :8080...")
//		err := http.ListenAndServe(":8080", nil)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	http.Handle("/", http.FileServer(http.Dir("./resources/")))

	port := ":8443"

	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	err := http.ListenAndServeTLS(port, "auth/cert.pem", "auth/key.pem", nil)

	log.Fatal(err)

}
