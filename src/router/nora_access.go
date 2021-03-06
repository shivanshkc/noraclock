package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"noraclock/src/constants"
)

func attachNoraAccess(router *mux.Router) *mux.Router {
	router.HandleFunc("/time", getTimeHandler).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/memory/{memoryID}", getMemoryHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/memory/{memoryID}", patchMemoryHandler).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/memory/{memoryID}", deleteMemoryHandler).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/memory", listMemoriesHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/memory", postMemoryHandler).Methods(http.MethodPost, http.MethodOptions)

	return router
}

func getTimeHandler(writer http.ResponseWriter, req *http.Request) {
	sendResponse(writer, http.StatusOK, nil, []byte(fmt.Sprintf(`{"time":%d}`, constants.NoraTime)))
}

func getMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

func patchMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

func deleteMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

func listMemoriesHandler(writer http.ResponseWriter, req *http.Request) {}

func postMemoryHandler(writer http.ResponseWriter, req *http.Request) {}
