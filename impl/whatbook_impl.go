package whatbookimpl

import (
	"github.com/labstack/echo/v4"
	api "whatbook.com/whatbook/api-library"
)

// WhatBook handlers implementation
type WhatBook struct {
}

// GetTest test endpoint
func (w WhatBook) GetTest(ctx echo.Context) error {
	return nil
}

// GetBooks book recommendation endpoint
func (w WhatBook) GetBooks(ctx echo.Context, params api.GetBooksParams) error {
	return nil
}
