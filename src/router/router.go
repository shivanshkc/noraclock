package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"noraclock/src/configs"
	"noraclock/src/logger"
	"noraclock/src/middleware"
	"os"
	"path"
	"path/filepath"
)

var conf = configs.Get()
var log = logger.General()

// Get returns the HTTP handler that is responsible to handle all API hits.
func Get() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.Interceptor)
	router.Use(middleware.CORS)

	apiRouter := attachAPIRouter(router.PathPrefix("/api").Subrouter())
	_ = attachNoraAccess(apiRouter.PathPrefix("/noraAccess").Subrouter())

	router.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		fileToServe := filepath.Join("public", path.Clean(req.URL.Path))
		fileInfo, err := os.Stat(fileToServe)

		if (os.IsNotExist(err) || !fileInfo.Mode().IsRegular()) && path.Ext(fileToServe) == "" {
			fileToServe = filepath.Join("public", "/index.html")
		}

		http.ServeFile(writer, req, fileToServe)
	})

	return router
}

func attachAPIRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("", func(writer http.ResponseWriter, req *http.Request) {
		resJSON := fmt.Sprintf(
			`{"name":"%s","version":"%s"}`,
			conf.Service.Name,
			conf.Service.Version,
		)
		sendResponse(writer, http.StatusOK, nil, []byte(resJSON))
	}).Methods(http.MethodGet, http.MethodOptions)

	return router
}
