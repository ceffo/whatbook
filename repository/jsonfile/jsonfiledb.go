package jsonfile

import (
	"encoding/json"
	"io/ioutil"

	api "whatbook.com/whatbook/api-library"
	"whatbook.com/whatbook/repository"
)

type JsonFileDb struct {
	Books []api.Book
}

// NewJSONFileDb reads the json file and returns a JsonFileDb
func NewJSONFileDb(fileName string) (repository.Db, error) {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var books []api.Book
	err = json.Unmarshal(content, &books)
	if err != nil {
		return nil, err
	}

	return &JsonFileDb{
		Books: books,
	}, nil
}

// GetAllBooks returns all books read from the json file
func (j *JsonFileDb) GetAllBooks() ([]api.Book, error) {

	return j.Books, nil
}
