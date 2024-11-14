package core

import (
	"github.com/inchestnov/bookmarks-cli/internal/browser"
	"github.com/inchestnov/bookmarks-cli/internal/core/source"
	"github.com/inchestnov/bookmarks-cli/internal/domain"
)

type ProcessOptions struct {
	Template string
	Engine   string
	Folder   string
	Mode     domain.ImportMode
}

func Process(src []byte, opts ProcessOptions) error {
	bookmarks, err := source.Parse(src, opts.Template)
	if err != nil {
		return err
	}

	folder := "default"
	if opts.Folder != "" {
		folder = opts.Folder
	}

	return browser.Import(bookmarks, opts.Engine, folder, opts.Mode)
}
