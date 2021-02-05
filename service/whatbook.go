package whatbook

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	api "whatbook.com/whatbook/api-library"
	"whatbook.com/whatbook/repository"
)

// Service handlers implementation
type Service struct {
	repo   repository.Db
	logger echo.Logger
}

// NewService instanciates a new endpoints handler
func NewService(db repository.Db, l echo.Logger) *Service {
	return &Service{
		repo:   db,
		logger: l,
	}
}

// GetTest test endpoint
func (w Service) GetTest(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// GetBooks book recommendation endpoint
func (w Service) GetBooks(ctx echo.Context, params api.GetBooksParams) error {

	ctx.Logger().Infoj(log.JSON{
		"endpoint": "GetBooks",
		"author":   params.Author,
		"genre":    params.Genre,
		"pages":    params.NumPages,
		"era":      params.Era,
	})
	books, err := w.repo.GetBooks(params)
	if err != nil {
		s := err.Error()
		errmsg := &api.Error{Message: &s}
		return ctx.JSON(http.StatusInternalServerError, errmsg)
	}

	return ctx.JSONPretty(http.StatusOK, books, "  ")
}
