package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":9090", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /snippet/view", snippetView)
	mux.HandleFunc("POST /snippet/create", snippetCreate)

	infoLog.Println("Starting server on", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
