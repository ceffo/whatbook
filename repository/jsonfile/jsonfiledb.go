package jsonfile

import (
	"encoding/json"
	"io/ioutil"

	api "whatbook.com/whatbook/api-library"
	"whatbook.com/whatbook/filtering"
	"whatbook.com/whatbook/repository"
)

// Db represents a fake database
type Db struct {
	Books []api.Book
}

// NewDb reads the json file and returns a JsonFileDb
func NewDb(fileName string) (repository.Db, error) {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var books []api.Book
	err = json.Unmarshal(content, &books)
	if err != nil {
		return nil, err
	}

	return &Db{
		Books: books,
	}, nil
}

// GetBooks retrieves books statisfing the given parameters
func (j *Db) GetBooks(apiFilter api.GetBooksParams) ([]api.Book, error) {

	filter := filtering.FromAPIFilter(apiFilter)
	var books []api.Book
	// apply filter
	for _, book := range j.Books {
		if filtering.FilterBook(&book, filter) {
			books = append(books, book)
		}
	}
	// sort
	filtering.SortBooks(books)

	return books, nil
}
