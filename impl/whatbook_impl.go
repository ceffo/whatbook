package whatbookimpl

import "github.com/labstack/echo/v4"

// WhatBook handlers implementation
type WhatBook struct {
}

// GetTest test endpoint
func (w WhatBook) GetTest(ctx echo.Context) error {
	return nil
}
