package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	// Imported because it is used in the new instance of a book

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
		//The variable book defines a slice of the data type called Book
		books, err := app.models.Books.GetAll()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//The code below calls the helper.go function to format, marshall, and write the json
		//the envelope that is wrapping the books variable is naming that collection of data books and then returning the data of the books variable
		if err := app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}
	//if the endpoint /v1/books is used with post, it does the following
	if r.Method == http.MethodPost {
		// fmt.Fprintln(w, "Added a new book to the reading list")
		//below are the pieces of information we expect that will then be unmarshalled into a go object
		//we are not using the Book struct that already exists because that contains different fields we don't need/want
		var input struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float64  `json:"rating"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// fmt.Fprintf(w, "%v\n", input) //this prints out the http response formatted with line breaks as the input struct

		book := &data.Book{
			Title:     input.Title,
			Published: input.Published,
			Pages:     input.Pages,
			Genres:    input.Genres,
			Rating:    float32(input.Rating),
		}

		err = app.models.Books.Insert(book)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("v1/books/%d", book.ID))

		//This writes the JSON response with a 201 Created status code and the Location header set
		err = app.writeJSON(w, http.StatusCreated, envelope{"book": book}, headers)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

}

// This is another Handler - an app method handling the get, update, deleting specific books
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
		return
	}

	//this will be removed when this application si connection to a database
	//this is using the struct from the internal/data package
	book, err := app.models.Books.Get(idInt)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	//The code below calls the helper.go function to format, marshall, and write the json
	//the envelope that is wrapping the book variable is naming that collection of data book and then returning the data of the book variable
	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// fmt.Fprintf(w, "Update the details of the book with ID: %d", idInt)

	book, err := app.models.Books.Get(idInt)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	//the struct below defines the way that we want to unmarshall the json
	//we are using pointers because we want to modify the existing struct instead of creating a new one

	var input struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float32 `json:"rating"`
	}

	// err = json.Unmarshal(body, &input) //not sure why we are doing this with err?

	err = app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Published != nil {
		book.Published = *input.Published
	}

	if input.Pages != nil {
		book.Pages = *input.Pages
	}

	//why doesn't this use an asterisk? Is it because we want it to be overwritten?
	if len(input.Genres) > 0 {
		book.Genres = input.Genres
	}

	if input.Rating != nil {
		book.Rating = *input.Rating
	}

	//why are we using the err variable for this?
	err = app.models.Books.Update(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete the book with ID: %d", idInt)
}
