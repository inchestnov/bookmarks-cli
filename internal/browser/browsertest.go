package browser

import (
	"github.com/inchestnov/bookmarks-cli/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrowser(t *testing.T, b Browser) {
	tests := []func(t *testing.T, b Browser){
		test_Browser_Replace_Then_GetAll,
	}

	for _, testFunc := range tests {
		testFunc(t, b)
	}
}

func test_Browser_Replace_Then_GetAll(t *testing.T, b Browser) {
	bookmarks, err := b.GetAll()
	assert.NoError(t, err)
	assert.Empty(t, bookmarks)

	expected := []domain.Bookmark{
		{
			Name: "1",
			URL:  "http://localhost:8080/1",
		},
		{
			Name: "2",
			Children: []domain.Bookmark{
				{
					Name: "21",
					URL:  "http://localhost:8080/21",
				},
				{
					Name: "22",
					URL:  "http://localhost:8080/21",
				},
			},
		},
	}
	assert.NoError(t, b.Replace(expected))

	bookmarks, err = b.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, bookmarks)
}
