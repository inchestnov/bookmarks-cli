package browser

import "github.com/inchestnov/bookmarks-cli/internal/domain"

type Browser interface {
	// GetAll get all stored bookmarks
	GetAll() ([]domain.Bookmark, error)

	// Replace all stored bookmarks
	Replace(bookmarks []domain.Bookmark) error
}
