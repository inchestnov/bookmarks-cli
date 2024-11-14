package domain

type Bookmark struct {
	Name     string
	URL      string
	Children []Bookmark
}
