package data

//Note: the internal directory carries special meaining and behavior in Go
//It means that any package under internal cannot be imported from outside this project

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// below is a struct that will be used to type a group of related data
// Without struct tags, the json will display exactly as the data is defined in the struct
// The struct tags added to the Book struct allow control over how the json is output when it is marshalled
type Book struct {
	ID        int64     `json:"id"` //this json tag changes the field name from ID to id
	CreatedAt time.Time `json:"-"`  //this json tag prevents this field from being displayed with the rest of the json when it is marshalled from the struct;
	//the above is in the database, but not displayed elsewhere after the json is marshalled
	Title     string   `json:"title"`               //this changes the title field to lower case
	Published int      `json:"published,omitempty"` //this json tag makes this field optional
	Pages     int      `json:"pages,omitempty"`
	Genres    []string `json:"genres,omitempty"`
	Rating    float32  `json:"rating,omitempty"`
	Version   int32    `json:"-"`
}

// this type is connected to all of the methods that implement the crud operations
type BookModel struct {
	DB *sql.DB
}

// this method "hangs off of" the BookModel type - like all of the following methods
// it takes in a pointer to a book - that is a pointer to a book record that is coming in to the database
func (b BookModel) Insert(book *Book) error {
	//the query variable holds the postgres sql statement that will be run to create a new record
	//the values are "positional arguments" and are being populated by the args variable below
	query := `
	INSERT INTO books (title, published, pages, genres, rating)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, version
	`

	//the blank interface below is taking in all the information from the pointer to a book above and then populates the query variable VALUES
	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating}
	//returns the auto-generated system values to Go object
	//this first runs the INSERT statement with the query and the args so the row is put into the database
	//it then returns back some values with the second part (which corresponds to the RETURNING part of the statement above)
	//the Scan part returns dereferenced pointers to those aspects of the book object because these are system generated
	return b.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

// this method takes in a book id and returns a pointer to a book and an error
func (b BookModel) Get(id int64) (*Book, error) {
	//this returns an error if the id is invalid
	if id < 1 {
		return nil, errors.New("record not found")
	}
	//this pulls the specific record from the database
	query := `
	SELECT id, created_at, title, published, pages, genres, rating, version
	FROM books
	WHERE id = $1
	`
	//this variable is used to hold all of the information for the book record from the database
	var book Book
	//Below passes back the scanned information
	//Scan is taking in the query and id information and then populating the variable with the record returned from the database
	err := b.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		pq.Array(&book.Genres),
		&book.Rating,
		&book.Version,
	)
	//this switch case is handling potential errors
	if err != nil {
		switch {
		//this case handles when there are no records with the specific id found
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err

		}
	}
	return &book, nil //this returns the book object with a nil error
}

func (b BookModel) Update(book *Book) error {
	query := `
	UPDATE books
	SET title = $1, published = $2, pages = $3, genres = $4, rating = $5, version = version +1
	WHERE id = $6 AND version = $7
	RETURNING version
	`

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating, book.ID, book.Version}
	return b.DB.QueryRow(query, args...).Scan(&book.Version)
}

func (b BookModel) Delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
	DELETE FROM books
	WHERE id = $1
	`

	results, err := b.DB.Exec(query, id)
	if err != nil {
		return err
	}

	//below checks that something actually happened by determining if any rows were changed in the database
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	//this returns an error if no rows were affected
	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

// GetAll doesn't take anything, but it does return a slice with pointers to books and an error
func (b BookModel) GetAll() ([]*Book, error) {
	query := `
	SELECT *
	FROM books
	ORDER BY id
	`

	rows, err := b.DB.Query(query)
	if err != nil {
		return nil, err
	}
	//the code below ends the database search when there are no more rows to find
	defer rows.Close()

	books := []*Book{} //this variable is a slice containing books of type Book

	//below convets the database record into an object
	//for each of the rows returned a book variable of type Book is created
	//then Scan in the information for that row into the book object

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.CreatedAt,
			&book.Title,
			&book.Published,
			&book.Pages,
			pq.Array(&book.Genres),
			&book.Rating,
			&book.Version,
		)
		if err != nil {
			return nil, err
		}

		//the book object is then added to the books variable
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
