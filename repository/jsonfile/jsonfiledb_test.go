package jsonfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fileName = "../../books/db.json"
)

func Test_ReadJsonFile(t *testing.T) {

	db, err := NewJSONFileDb(fileName)
	if assert.NoError(t, err, "Reading file must not return an error") {

		books, err := db.GetAllBooks()

		if assert.NoError(t, err, "GetAllBooks must not return an error") {
			assert.NotZero(t, len(books), "No books returned")
		}
	}
}
