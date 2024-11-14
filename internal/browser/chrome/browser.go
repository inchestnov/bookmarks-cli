package chrome

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/inchestnov/bookmarks-cli/internal/browser"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var _ browser.Browser = &Browser{}

type Browser struct {
	profilePath string
}

func NewProfileBrowser(location, profileName string) *Browser {
	profilePath := filepath.Join(location, profileName)
	return &Browser{
		profilePath: profilePath,
	}
}

func NewDefaultBrowser() *Browser {
	// FIXME 2025-01-23: For local testing
	return NewProfileBrowser("/Users/inchestnov/Library/Application Support/Google/Chrome", "Profile 1")
}

func (b *Browser) GetAll() ([]domain.Bookmark, error) {
	root, err := b.findRoot()
	if err != nil {
		return nil, err
	}

	return convertToDomain(root.Children), nil
}

func (b *Browser) Replace(bookmarks []domain.Bookmark) error {
	root, err := b.findRoot()
	if err != nil {
		return err
	}

	// No root exists yet
	if root.Name == "" {
		root = Bookmark{
			Type:         TypeFolder,
			Name:         "Bookmarks CLI (Imported)",
			Guid:         uuid.NewString(), // FIXME 2024-12-11: ?
			ID:           "",               // FIXME 2024-12-11: ?
			DateLastUsed: "0",              // FIXME 2024-12-11: ?
			DateAdded:    strconv.FormatInt(time.Now().Unix(), 10),
			DateModified: strconv.FormatInt(time.Now().Unix(), 10),
		}
	}

	converted := convertFromDomain(bookmarks)
	root.Children = converted

	return b.updateRoot(root)
}

func (b *Browser) findRoot() (Bookmark, error) {
	bookmarksFileContent, err := os.ReadFile(filepath.Join(b.profilePath, "Bookmarks"))
	if err != nil {
		return Bookmark{}, err
	}

	var bookmarksFile BookmarksFile
	if err = json.Unmarshal(bookmarksFileContent, &bookmarksFile); err != nil {
		return Bookmark{}, err
	}

	for _, v := range bookmarksFile.Roots.Other.Children {
		if v.Name == "Bookmarks CLI (Imported)" {
			return v, nil
		}
	}

	return Bookmark{}, nil
}

func (b *Browser) updateRoot(root Bookmark) error {
	bookmarksFileContent, err := os.ReadFile(filepath.Join(b.profilePath, "Bookmarks"))
	if err != nil {
		return err
	}

	var bookmarksFile BookmarksFile
	if err = json.Unmarshal(bookmarksFileContent, &bookmarksFile); err != nil {
		return err
	}

	rootIndex := -1
	for i, v := range bookmarksFile.Roots.Other.Children {
		if v.Name == "Bookmarks CLI (Imported)" {
			rootIndex = i
			break
		}
	}

	if rootIndex == -1 {
		bookmarksFile.Roots.Other.Children = append(bookmarksFile.Roots.Other.Children, root)
	} else {
		bookmarksFile.Roots.Other.Children[rootIndex] = root
	}

	bookmarksFileContent, err = json.Marshal(bookmarksFile)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(b.profilePath, "Bookmarks"), bookmarksFileContent, 0644)
}
