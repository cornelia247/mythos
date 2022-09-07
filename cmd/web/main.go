package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/cornelia247/mythos/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	myths    *mysql.MythModel
}

func main() {
	//address variable to listen on, flags also created
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:mypassword@/mythos?parseTime=true", "MySQL data source name")
	flag.Parse()

	// error and info types incase of any problem when compling/running application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	//Creating a database connnection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)

	}
	defer db.Close()
	// Initialize a mysql.MythModel instance and add it to the application // dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		myths:    &mysql.MythModel{DB: db},
	}

	// creating our custom listen and serve and starting the server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting Server On %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log   incase you want to log to a file
