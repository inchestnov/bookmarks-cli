package chrome

import (
	"github.com/google/uuid"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
	"strconv"
	"time"
)

func convertToDomain(bookmarks []Bookmark) []domain.Bookmark {
	result := make([]domain.Bookmark, 0, len(bookmarks))
	for _, b := range bookmarks {
		d := domain.Bookmark{
			Name: b.Name,
		}

		if b.Type == TypeFolder {
			children := convertToDomain(b.Children)
			d.Children = children
		} else if b.Type == TypeURL {
			d.URL = b.URL
		} else {
			// unreachable
			panic("unknown bookmark type: " + b.Type)
		}

		result = append(result, d)
	}
	return result
}

func convertFromDomain(bookmarks []domain.Bookmark) []Bookmark {
	result := make([]Bookmark, 0, len(bookmarks))
	for _, d := range bookmarks {
		b := Bookmark{
			Name:         d.Name,
			Guid:         uuid.NewString(), // FIXME 2024-12-11: ?
			ID:           "",               // FIXME 2024-12-11: ?
			DateLastUsed: "0",              // FIXME 2024-12-11: ?
			DateAdded:    strconv.FormatInt(time.Now().Unix(), 10),
			DateModified: strconv.FormatInt(time.Now().Unix(), 10),
		}

		if d.URL == "" {
			b.Type = TypeFolder
			children := convertFromDomain(d.Children)
			b.Children = children
		} else {
			b.Type = TypeURL
			b.URL = d.URL
		}
		result = append(result, b)
	}
	return result
}

//func convert(groupName string, bookmark []domain.Bookmark) Bookmark {
//	groups := lo.GroupBy(bookmark, func(item domain.Bookmark) string {
//		return item.Group
//	})
//
//	group := Bookmark{
//		Type:         TypeFolder,
//		Name:         groupName,
//		Guid:         uuid.NewString(),
//		ID:           "", // fixme ???
//		DateAdded:    strconv.FormatInt(time.Now().Unix(), 10),
//		DateLastUsed: "0",
//		DateModified: strconv.FormatInt(time.Now().Unix(), 10),
//		Children:     make([]Bookmark, 0, len(groups)),
//	}
//	for g, items := range groups {
//		converted := convertGroup(g, items)
//		group.Children = append(group.Children, converted)
//	}
//	return group
//}
//
//func convertGroup(name string, items []domain.Bookmark) Bookmark {
//	group := Bookmark{
//		Type:         TypeFolder,
//		Name:         name,
//		Guid:         uuid.NewString(),
//		ID:           "", // fixme ???
//		DateAdded:    strconv.FormatInt(time.Now().Unix(), 10),
//		DateLastUsed: "0",
//		DateModified: strconv.FormatInt(time.Now().Unix(), 10),
//		Children:     make([]Bookmark, 0, len(items)),
//	}
//
//	for _, item := range items {
//		group.Children = append(group.Children, Bookmark{
//			Type:         TypeURL,
//			Name:         item.Name,
//			URL:          item.URL,
//			Guid:         uuid.NewString(),
//			ID:           "", // fixme ???
//			DateAdded:    strconv.FormatInt(time.Now().Unix(), 10),
//			DateLastUsed: "0",
//			DateModified: strconv.FormatInt(time.Now().Unix(), 10),
//		})
//	}
//
//	return group
//}
