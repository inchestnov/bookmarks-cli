package source

import (
	"github.com/inchestnov/bookmarks-cli/internal/core/source/regular"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
)

const (
	templateRegular = "regular"
)

type Parser interface {
	Parse([]byte) ([]domain.Bookmark, error)
}

func Parse(content []byte, template string) ([]domain.Bookmark, error) {
	var p Parser
	switch template {
	default:
		p = regular.New()
	case templateRegular:
		p = regular.New()
	}

	return p.Parse(content)
}
