package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time" // Imported because it is used in the new instance of a book

	"readinglist/internal/data" // this imports the data package; one can use the cat go.mod command in terminal to determine how to begin import statement if needed
)

// app method handling healthcheck endpoint
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//the encoding/marshalling for the healthcheck endpoint will be done differently from the others
	//it's not going to use a struct to convert the json to and from the messages, it's going to use native types
	//it's going to assume based on the data type of the go object itself what type of json values should be marshalled into the response

	//using the data variable is expected
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	//below turns the data map from above into json
	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return // this exits, stopping the rest of the code from running
	}
	//This formats the json some
	js = append(js, '\n')

	//Below sets the http headers
	w.Header().Set("Content-Type", "application/json")

	//Below write the http response - we pass in the json object and that is what will be written in the response
	w.Write(js)
}

// This is a Handler - an app method handling getting and creating new books within the total list of books
func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	//the if statement validates that the request at this endpoint is only either GET or POST

	//if the endpoint /v1/books is used with get, it does the following
	if r.Method == http.MethodGet {
		// fmt.Fprintln(w, "Display a list of books on the reading list")

		//The variable book defines a slice of the data type called Book
		books := []data.Book{
			{
				ID:        1,
				CreatedAt: time.Now(),
				Title:     "The Killer Called Collect",
				Published: 1985,
				Pages:     195,
				Genres:    []string{"Fiction", "Thriller", "Mystery"},
				Rating:    4.7,
				Version:   1,
			},
			{
				ID:        2,
				CreatedAt: time.Now(),
				Title:     "A Lilliard Family History",
				Published: 1990,
				Pages:     490,
				Genres:    []string{"Non-Fiction", "Historical"},
				Rating:    2.8,
				Version:   1,
			},
		}

		js, err := json.MarshalIndent(books, "", "\t") //This version of marshalling indents the displayed json and keys - in this case with a tab
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		js = append(js, '\n')

		w.Header().Set("Content-Type", "application/json")

		w.Write(js)
		return
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
		app.updateBook(w, r)

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
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Display the details of book with ID: %d", idInt)
	//for now, hard-coding a new instance of the book object
	//this will be removed when this application si connection to a database
	//this is using the struct from the internal/data package
	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "The Corpse Danced at Midnight",
		Published: 1984,
		Pages:     190,
		Genres:    []string{"Fiction", "Thriller", "Mystery"},
		Rating:    4.8,
		Version:   1,
	}

	//Below creates the json object by marshalling the new instance of the Book struct
	//because the book variable is a struct, Go knows how to create the json object with the correct data types
	js, err := json.Marshal(book)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update the details of the book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete the book with ID: %d", idInt)
}
