package chrome

import (
	_ "embed"
	"github.com/inchestnov/bookmarks-cli/internal/browser"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

//go:embed bookmarks.example.json
var bookmarks []byte

func TestBrowser(t *testing.T) {
	tempDir := os.TempDir()

	root, err := os.MkdirTemp(tempDir, "chrome_root")
	assert.NoError(t, err)

	assert.NoError(t, os.MkdirAll(filepath.Join(root, "JUnit"), 0777))

	assert.NoError(t, os.WriteFile(filepath.Join(root, "JUnit", "Bookmarks"), bookmarks, 0644))

	b := NewProfileBrowser(root, "JUnit")

	browser.TestBrowser(t, b)
}
