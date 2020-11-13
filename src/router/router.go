package router

import (
	"fmt"
	"net/http"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/constants"
	"noraclock/v2/src/middleware"

	"github.com/gorilla/mux"
)

// Handlers : Returns the HTTP Handlers.
func Handlers() http.Handler {
	router := mux.NewRouter()

	router.Use(middleware.Interceptor)
	router.Use(middleware.CORS)
	router.Use(middleware.NoraGuard)
	router.Use(middleware.ResponseHeader)

	router.HandleFunc("/api/noraAccess/memory/{memoryID}", getMemoryByIDHandler).
		Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/noraAccess/memory/{memoryID}", patchMemoryHandler).
		Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/api/noraAccess/memory/{memoryID}", deleteMemoryHandler).
		Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/api/noraAccess/memory", getMemoryHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/api/noraAccess/memory", postMemoryHandler).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/noraAccess/time", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"time":%d}`, constants.NoraTime)))
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(fmt.Sprintf("Hi, I'm %s. %s.", configs.Service.Name, configs.Service.Description)))
	}).Methods(http.MethodGet)

	return router
}
