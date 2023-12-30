package main

import (
	"flag"
	"log"
	"net/http"

	"readinglist/internal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")                                                        //using a flag means we can change it with command line arguments
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/books", "Endpoint for the readinglist web service") //this defines the endpoint for the readinglist

	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint}, //this sets the endpoint as a pointer to the endpoint flag that was passed in
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe() //this starts the web application
	log.Fatal(err)
}
