package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/myth", showMyth)
	mux.HandleFunc("/myth/create", createMyth)

	log.Println("Starting Server On :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
