package filtering

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	api "whatbook.com/whatbook/api-library"
)

func TestYearBelowThresholdIsClassical(t *testing.T) {
	year := modernYearThreshold - 1
	era := yearToEra(year)
	assert.Equal(t, era, Classical)
}

func TestYearEqualToThresholdIsClassical(t *testing.T) {
	year := modernYearThreshold
	era := yearToEra(year)
	assert.Equal(t, era, Classical)
}

func TestYearGreaterThanThresholdIsModern(t *testing.T) {
	year := modernYearThreshold + 1
	era := yearToEra(year)
	assert.Equal(t, era, Modern)
}

func TestFilterBook(t *testing.T) {
	type args struct {
		book   *api.Book
		filter Filter
	}

	author1 := "Zaphod BeebleBrox"
	author2 := "Ford Prefect"
	genre1 := "Fictional"
	genre2 := "Fantastic"
	var modernYear uint32 = 1960
	var classicalYear uint32 = 1750
	var numpages1 uint32 = 200
	var numpages2 uint32 = 300

	book1 := api.Book{
		Author:   author1,
		Genre:    genre1,
		NumPages: numpages1,
		Year:     modernYear,
	}

	book2 := api.Book{
		Author:   author2,
		Genre:    genre2,
		NumPages: numpages2,
		Year:     classicalYear,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty Filter matches",
			args: args{
				book:   &book1,
				filter: Filter{},
			},
			want: true,
		},
		{
			name: "Author matches",
			args: args{
				book: &book1,
				filter: Filter{
					Author: &author1,
				},
			},
			want: true,
		},
		{
			name: "Author mismatch",
			args: args{
				book: &book1,
				filter: Filter{
					Author: &author2,
				},
			},
			want: false,
		},
		{
			name: "Genre matches",
			args: args{
				book: &book1,
				filter: Filter{
					Genre: &genre1,
				},
			},
			want: true,
		},
		{
			name: "Genre mismatch",
			args: args{
				book: &book1,
				filter: Filter{
					Genre: &genre2,
				},
			},
			want: false,
		},
		{
			name: "Num pages matches",
			args: args{
				book: &book1,
				filter: Filter{
					NumPages: &numpages1,
				},
			},
			want: true,
		},
		{
			name: "Num pages mismatch",
			args: args{
				book: &book1,
				filter: Filter{
					NumPages: &numpages2,
				},
			},
			want: false,
		},
		{
			name: "Modern year matches",
			args: args{
				book: &book1,
				filter: Filter{
					Era: Modern,
				},
			},
			want: true,
		},
		{
			name: "Modern year mismatch",
			args: args{
				book: &book1,
				filter: Filter{
					Era: Classical,
				},
			},
			want: false,
		},
		{
			name: "Classicl year matches",
			args: args{
				book: &book2,
				filter: Filter{
					Era: Classical,
				},
			},
			want: true,
		},
		{
			name: "Classical year mismatch",
			args: args{
				book: &book2,
				filter: Filter{
					Era: Modern,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterBook(tt.args.book, tt.args.filter); got != tt.want {
				t.Errorf("FilterBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortBooks(t *testing.T) {

	var books []api.Book
	numBooks := 20
	for i := 0; i < numBooks; i++ {
		books = append(books, api.Book{
			Rating: 1 + rand.Uint32()%5,
		})
	}

	SortBooks(books)

	assert.True(t, sort.SliceIsSorted(books, func(i, j int) bool {
		return books[i].Rating > books[j].Rating
	}), "Books are not sorted by decreasing rating")
}
