package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// creating a new server their handlers for various routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/myth", app.showMyth)
	mux.HandleFunc("/myth/create", app.createMyth)

	// serving content
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
