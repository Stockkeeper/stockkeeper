package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed build
var embeddedFS embed.FS

// GetHttpFileSystem returns a http.FileSystem that contains static files for
// the frontend.
func GetHttpFileSystem() http.FileSystem {
	subFS, err := fs.Sub(embeddedFS, "build")
	if err != nil {
		panic(err)
	}
	return http.FS(subFS)
}

func GetFileSystem() fs.FS {
	subFS, err := fs.Sub(embeddedFS, "build")
	if err != nil {
		panic(err)
	}
	return subFS
}
