package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:build/web
var embedFrontend embed.FS

// GetFileSystem returns the subtree of the embedded files
// so we don't have to deal with the "build/web" prefix in our routes.
func GetFileSystem() fs.FS {
	f, err := fs.Sub(embedFrontend, "build/web")
	if err != nil {
		panic(err)
	}
	return f
}
