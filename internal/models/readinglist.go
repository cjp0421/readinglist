package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// the 3 types below allow us to unmarshall json
type Book struct { //type for each book in the envelopes
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float32  `json:"rating"`
}

type BookResponse struct { //type for enveloped single-book json responses
	Book *Book `json:"book"` //pointer to the book
}

type BooksResponse struct { //type for enveloped multi-book json responses
	Books *[]Book `json:"books"` //pointer to a slice of books
}

type ReadinglistModel struct { //this type is what all of the methods "hang on to"
	Endpoint string //this is the url to the web service
}

// the method below returns all of the book records in the database for the homepage
// it doesn't take any values but it returns a slice of books and an error
func (m *ReadinglistModel) GetAll() (*[]Book, error) { //it is a method that hangs off of the dereferenced pointer to ReadinglistModel
	resp, err := http.Get(m.Endpoint) //this is what passes the url into the web service
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body) //this reads the response body and puts it into a variable called data
	if err != nil {
		return nil, err
	}

	var booksResp BooksResponse //this handles the envelope

	err = json.Unmarshal(data, &booksResp) //this unmarshalls the response in the data variable and puts it into the booksResp variable
	if err != nil {
		return nil, err
	}

	return booksResp.Books, nil //this returns the specific books data inside the envelope (does not return the envelope)
}

// this method takes in an id and returns a pointer to a book and an error - it returns a specific book by id
func (m *ReadinglistModel) Get(id int64) (*Book, error) {
	url := fmt.Sprintf("%s/%d", m.Endpoint, id) //this makes the url variable contain a string with the endpoint and the id; it formats it fit the url style
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bookResp BookResponse //this creates a variable of type BookResponse (singular) to deal with then envelope

	err = json.Unmarshal(data, &bookResp)
	if err != nil {
		return nil, err
	}

	return bookResp.Book, nil //this returns the singular book without the envelope
}
