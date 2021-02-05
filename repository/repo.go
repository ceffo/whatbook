package repository

import (
	api "whatbook.com/whatbook/api-library"
)

// Db represents the interface to the data repo
type Db interface {
	GetBooks(filter api.GetBooksParams) ([]api.Book, error)
}
