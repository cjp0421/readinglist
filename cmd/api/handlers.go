package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// app method handling healthcheck endpoint
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

// This is a Handler - an app method handling getting and creating new books within the total list of books
func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	//the if statement validates that the request at this endpoint is only either GET or POST

	//if the endpoint /v1/books is used with get, it does the following
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of books on the reading list")
	}
	//if the endpoint /v1/books is used with post, it does the following
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new book to the reading list")
	}
}

// This is another Handler (starts at 745) - an app method handling the get, update, deleting specific books
// Below is a request multiplexer (aka a request router). It routes incoming requests to a handler using a set of rules
func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)

	case http.MethodPut:
		app.udpateBook(w, r)

	case http.MethodDelete:
		app.deleteBook(w, r)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//Below is the definition of each specific case above

// each of the methods below need to have a way to get the id of the book in question from the URL
// getting a specific book
func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books/"):]
	idInt := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
