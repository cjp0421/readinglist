package data

import "database/sql"

//this file is intended to encapsulate the different models being used

type Models struct {
	Books BookModel
}

// the function below just returns the model
// it takes in a pointer to a SQL database
// this helps us connect to the database and then implement CRUD operations
func NewModels(db *sql.DB) Models {
	return Models{
		Books: BookModel{DB: db},
	}
}
