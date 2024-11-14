package _import

import (
	"github.com/inchestnov/bookmarks-cli/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "embed"
)

//go:embed regular.example.json
var regular []byte

func TestImportRegular(t *testing.T) {
	bookmarks, err := ImportRegular(regular)
	assert.NoError(t, err)

	assert.Equal(t, []domain.Bookmark{
		{
			Name: "First",
			Children: []domain.Bookmark{
				{
					Name: "/11",
					URL:  "http://localhost:8181/11",
				},
				{
					Name: "/12",
					URL:  "http://localhost:8182/12",
				},
			},
		},
		{
			Name: "Second",
			Children: []domain.Bookmark{
				{
					Name: "/21",
					URL:  "http://localhost:8281/21",
				},
			},
		},
		{
			Name: "/3",
			URL:  "http://localhost:8380/3",
		},
	}, bookmarks)
}
