package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "hl1-proxy/cmd/web/docs"
)

// @title HL-Task1 Proxy Swagger
// @version 1.0
// @description swagger for proxy-server
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", ":4000", "HTTP network address")
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}
	infoLog.Printf("Starting server on %+s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
