package main

import (
	"log"
	"net/http"
)

const listenAddr = ":9090"

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /snippet/view", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	log.Println("Starting server on", listenAddr)
	err := http.ListenAndServe(listenAddr, mux)
	log.Fatal(err)
}
