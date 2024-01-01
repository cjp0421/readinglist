package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// these functions are all methods on the application type
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //tests to make sure the path the request is on is / to access - ensures that visitors will land on the homepage
		http.NotFound(w, r)
		return
	}

	books, err := app.readinglist.GetAll() //populating variable books with all of the book records in the database and return them as a Go object
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//below is a variable that is a slice of strings that have the path to the templates we want to use on the home page
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	//this parses the files at the string locations in the variable files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//this executes the template base first and then pulls in other templates
	err = ts.ExecuteTemplate(w, "base", books) //this takes in the io writer, the name of the first template, and the data contained in books
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

}

func (app *application) bookView(w http.ResponseWriter, r *http.Request) { //this returns a single book record
	id, err := strconv.Atoi(r.URL.Query().Get("id")) //this gets the id from the URL and converts it from a string to an int
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	book, err := app.readinglist.Get(int64(id)) //this get the specific book linked to the id int converted to the int64 type
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s (%d)\n", book.Title, book.Pages) //prints out the title of the specific book and the number of pages
}

// the method below needs to use both the GET method and the POST method
// GET to display the form and POST to update the database with the new book record
// because we need use two methods, we will use a mux (multiplexer)
func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.bookCreateForm(w, r)
	case http.MethodPost:
		app.bookCreateProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) bookCreateForm(w http.ResponseWriter, r *http.Request) {
	//below the html for the form is written
	fmt.Fprintf(w, "<html><head><title>Create Book</title></head>"+
		"<body><h1>Create Book</h1><form action=\"/book/create\" method=\"post\">"+
		"<label for=\"title\">Title</label><input type=\"text\" name=\"title\" id=\"title\">"+
		"<label for=\"pages\">Pages</label><input type=\"number\" name=\"pages\" id=\"pages\">"+
		"<label for=\"published\">Published</label><input type=\"number\" name=\"published\" id=\"published\">"+
		"<label for=\"genres\">Genres</label><input type=\"text\" name=\"genres\" id=\"genres\">"+
		"<label for=\"rating\">Rating</label><input type=\"number\" step=\"0.1\" name=\"rating\" id=\"rating\">"+
		"<button type=\"submit\">Create</form></body></html>")
}

func (app *application) bookCreateProcess(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title") //grabs the title from the POST request
	if title == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	pages, err := strconv.Atoi(r.PostFormValue("pages"))
	if err != nil || pages < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	published, err := strconv.Atoi(r.PostFormValue("published"))
	if err != nil || published < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	genres := strings.Split(r.PostFormValue("genres"), " ")

	//deviated from video and made this an int due to issues getting it to work when using a float64/float32 in some places
	ratingFloat, err := strconv.ParseFloat(r.PostFormValue("rating"), 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rating := float32(ratingFloat)

	//below is an anonymous struct - those have to be instantiated right away
	book := struct {
		Title     string   `json:"title"`
		Pages     int      `json:"pages"`
		Published int      `json:"published"`
		Genres    []string `json:"genres"`
		Rating    float32  `json:"rating"`
	}{ //this is a struct literal
		Title:     title,
		Pages:     pages,
		Published: published,
		Genres:    genres,
		Rating:    rating,
	}

	data, err := json.Marshal(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req, _ := http.NewRequest("POST", app.readinglist.Endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	//below creates the client
	client := &http.Client{}
	resp, err := client.Do(req) //this is where the request is actually sent to the web service
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("unexpected status: %s", resp.Status)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther) //redirects to the homepage which will be updated with the new book as landing there sends a request for all books in the table
}
