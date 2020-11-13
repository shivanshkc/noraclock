package router

import (
	"net/http"
	"noraclock/v2/src/exception"
	"os"
	"path/filepath"
)

type static struct {
	staticPath string
	indexPath  string
}

var staticHandler = &static{"static", "index.html"}

func (s *static) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	// Get the absolute path to prevent directory traversal.
	path, err := filepath.Abs(req.URL.Path)
	if err != nil {
		sendError(writer, exception.BadRequest(""))
		return
	}

  // Prepend the path with the path to the static directory.
	path = filepath.Join(s.staticPath, path)

  // Check whether a file exists at the given path.
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		http.ServeFile(writer, req, filepath.Join(s.staticPath, s.indexPath))
		return
	} else if err != nil {
		// If we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 and stop.
		sendError(writer, exception.Unexpected(""))
		return
	}

	// Otherwise, use http.FileServer to serve the static dir.
	http.FileServer(http.Dir(s.staticPath)).ServeHTTP(writer, req)
}