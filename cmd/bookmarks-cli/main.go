package main

import (
	"github.com/inchestnov/bookmarks-cli/internal/browser/chrome"
	_import "github.com/inchestnov/bookmarks-cli/internal/import"
	"log"
	"os"
)

func main() {
	chromeBrowser := chrome.NewDefaultBrowser()

	regularContent, err := os.ReadFile("/Users/inchestnov/go/projects/bookmarks-cli/data/regular.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = _import.ImportRegularIntoBrowser(regularContent, chromeBrowser); err != nil {
		log.Fatal(err)
	}
}
