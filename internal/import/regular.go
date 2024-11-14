package _import

import (
	"encoding/json"
	"github.com/inchestnov/bookmarks-cli/internal/browser"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
)

type RegularItem struct {
	Name     string        `json:"name"`
	URL      string        `json:"url"`
	Children []RegularItem `json:"children"`
}

func ImportRegular(content []byte) ([]domain.Bookmark, error) {
	var items []RegularItem
	if err := json.Unmarshal(content, &items); err != nil {
		return nil, err
	}

	return convertToDomain(items), nil
}

func ImportRegularIntoBrowser(content []byte, b browser.Browser) error {
	imported, err := ImportRegular(content)
	if err != nil {
		return err
	}

	return b.Replace(imported)
}

func convertToDomain(items []RegularItem) []domain.Bookmark {
	result := make([]domain.Bookmark, 0, len(items))
	for _, b := range items {
		d := domain.Bookmark{
			Name: b.Name,
		}

		if len(b.Children) > 0 {
			children := convertToDomain(b.Children)
			d.Children = children
		} else {
			d.URL = b.URL
		}

		result = append(result, d)
	}
	return result
}
