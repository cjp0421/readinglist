package data

//Note: the internal directory carries special meaining and behavior in Go
//It means that any package under internal cannot be imported from outside this project

import (
	"time"
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
