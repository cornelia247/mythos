package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//address variable to listen on, flags also created
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// error and info types incase of any problem when compling/running application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// creating our custom listen and serve and starting the server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting Server On %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log   incase you want to log to a file
