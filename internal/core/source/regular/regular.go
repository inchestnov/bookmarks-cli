package regular

import (
	"encoding/json"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
)

type Item struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Children []Item `json:"children"`
}

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (r *Parser) Parse(content []byte) ([]domain.Bookmark, error) {
	var root []Item
	if err := json.Unmarshal(content, &root); err != nil {
		return nil, err
	}

	return convert(root), nil
}

func convert(items []Item) []domain.Bookmark {
	converted := make([]domain.Bookmark, 0, len(items))
	for _, v := range items {
		d := domain.Bookmark{
			Name: v.Name,
		}

		if v.URL == "" {
			children := convert(v.Children)
			d.Children = children
		} else {
			d.URL = v.URL
		}
		converted = append(converted, d)
	}

	return converted
}
