package jsonfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	api "whatbook.com/whatbook/api-library"
)

const (
	fileName = "../../books/db.json"
)

func TestReadJsonFile(t *testing.T) {

	db, err := NewDb(fileName)
	if assert.NoError(t, err, "Reading file must not return an error") {

		books, err := db.GetBooks(api.GetBooksParams{})

		if assert.NoError(t, err, "GetBooks must not return an error") {
			assert.NotZero(t, len(books), "No books returned")
		}
	}
}
