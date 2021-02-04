package repository

import (
	api "whatbook.com/whatbook/api-library"
)

type Db interface {
	GetAllBooks() ([]api.Book, error)
}
