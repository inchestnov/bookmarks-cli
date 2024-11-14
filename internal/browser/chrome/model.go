package chrome

const (
	TypeFolder = "folder"
	TypeURL    = "url"
)

type Bookmark struct {
	Children     []Bookmark `json:"children"`
	DateAdded    string     `json:"date_added"`
	DateLastUsed string     `json:"date_last_used"`
	DateModified string     `json:"date_modified"`
	Guid         string     `json:"guid"`
	ID           string     `json:"id"`
	MetaInfo     MetaInfo   `json:"meta_info,omitempty"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	URL          string     `json:"url,omitempty"`
}

type MetaInfo struct {
	PowerBookmarkMeta string `json:"power_bookmark_meta"`
}

type Roots struct {
	BookmarkBar Bookmark `json:"bookmark_bar"`
	Other       Bookmark `json:"other"`
	Synced      Bookmark `json:"synced"`
}

type BookmarksFile struct {
	Checksum string `json:"checksum"`
	Roots    Roots  `json:"roots"`
	Version  int    `json:"version"`
}
