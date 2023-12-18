package data

//Note: the internal directory carries special meaining and behavior in Go
//It means that any package under internal cannot be imported from outside this project

import (
	"time"
)

// below is a struct that will be used to type a group of related data
type Book struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Published int
	Pages     int
	Genres    []string
	Rating    float32
	Version   int32
}
