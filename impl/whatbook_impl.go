package whatbookimpl

import (
	"net/http"

	"github.com/labstack/echo/v4"
	api "whatbook.com/whatbook/api-library"
	"whatbook.com/whatbook/repository"
)

// WhatBook handlers implementation
type WhatBook struct {
	repo repository.Db
}

func NewWhatBookAPI(db repository.Db) *WhatBook {
	return &WhatBook{
		repo: db,
	}
}

// GetTest test endpoint
func (w WhatBook) GetTest(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// GetBooks book recommendation endpoint
func (w WhatBook) GetBooks(ctx echo.Context, params api.GetBooksParams) error {

	books, err := w.repo.GetAllBooks()
	if err != nil {
		s := err.Error()
		errmsg := &api.Error{Message: &s}
		return ctx.JSON(http.StatusInternalServerError, errmsg)
	}

	return ctx.JSONPretty(http.StatusOK, books, "  ")
}
