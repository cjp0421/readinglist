package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address") //using a flag means we can change it with command line arguments

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe() //this starts the web application
	log.Fatal(err)
}
