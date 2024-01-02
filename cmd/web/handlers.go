package main

import (
	"bytes"
	"encoding/json"
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

	//this pulls in the files with the templates we'll combine to create the page
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/view.html",
	}

	//used to convert comma-separated genres to a slice within the template
	funcs := template.FuncMap{"join": strings.Join}

	ts, err := template.New("showBook").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
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
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/create.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// nil is used here because we don't have any data we want to pass in
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Sever Error", 500)
		return
	}
}

func (app *application) bookCreateProcess(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//the value in the Get method needs to match the name attribute in the html form
	title := r.PostForm.Get("title")

	//published and pages needs to include error handling because they are being converted from one data type to another
	published, err := strconv.Atoi(r.PostForm.Get("published"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	pages, err := strconv.Atoi(r.PostForm.Get("pages"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//genres is being split apart at each comma and converting it into a slice of strings
	genres := strings.Split(r.PostForm.Get("genres"), ",")

	rating, err := strconv.ParseFloat(r.PostForm.Get("rating"), 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

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
		Rating:    float32(rating),
	}

	data, err := json.Marshal(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", app.readinglist.Endpoint, bytes.NewBuffer(data))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

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
