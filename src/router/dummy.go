package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func attachDummyRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("", postDummyHandler).Methods(http.MethodPost)
	return router
}

func postDummyHandler(writer http.ResponseWriter, req *http.Request) {
	sendResponse(writer, http.StatusCreated, nil, nil)
}
