package main

import "net/http"

//this is a method tied to application that returns a new ServeMux
func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck) // this is an route
	// Endpoints are functions available through the API
	// A route is the name you use to access endpoints, used in the URL
	
	mux.HandleFunc("/v1/books", app.getCreateBooksHandler) // Gets all books with the GET method, Creates new book with the POST method
	//1st arg is the route; 2nd arg is the handler function (endpoint)
	
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler) // Handles queries related to individual books
	mux.HandleFunc
}