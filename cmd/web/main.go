package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//address variable to listen on, flags also created
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// error and info types incase of any problem when compling/running application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// creating a new server their handlers for various routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/myth", showMyth)
	mux.HandleFunc("/myth/create", createMyth)

	// serving content
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// starting the server
	infoLog.Printf("Starting Server On %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
