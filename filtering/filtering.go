package filtering

import (
	"sort"
	"strings"

	api "whatbook.com/whatbook/api-library"
)

// Era represents an era enumeration
type Era int

const (
	// Dontcare era
	Dontcare Era = iota
	// Classical era
	Classical
	// Modern era
	Modern
)

const (
	// threshold to be considered modern (stricly greater)
	modernYearThreshold uint32 = 1850
)

// Filter holds the different criteria for filtering books
type Filter struct {
	Author   *string
	Genre    *string
	NumPages *uint32
	Era      Era
}

// FromAPIFilter converts the api filter
func FromAPIFilter(apiFilter api.GetBooksParams) Filter {
	return Filter{
		Author: apiFilter.Author,
	}
}

// getEra gets the era of a given year
func getEra(year uint32) Era {
	if year > modernYearThreshold {
		return Modern
	}
	return Classical
}

// FilterBook returns whether a given book satisfies the provided filter
func FilterBook(book *api.Book, filter Filter) bool {
	if !filterOptionalString(filter.Author, book.Author) {
		return false
	}

	if !filterOptionalString(filter.Genre, book.Genre) {
		return false
	}

	if !filterOptionalUint32(filter.NumPages, book.NumPages) {
		return false
	}

	era := getEra(book.Year)
	if filter.Era != Dontcare && filter.Era != era {
		return false
	}

	return true
}

// SortBooks sorts a Book slice by decreasing rating
func SortBooks(books []api.Book) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Rating > books[j].Rating
	})
}

func filterOptionalString(filter *string, s string) bool {
	return filter == nil || strings.EqualFold(*filter, s)
}

func filterOptionalUint32(filter *uint32, u uint32) bool {
	return filter == nil || *filter == u
}
