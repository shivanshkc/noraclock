package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"noraclock/src/business"
	"noraclock/src/constants"
	"noraclock/src/exception"
	"noraclock/src/middleware"
	"noraclock/src/validator"
)

func attachNoraAccess(router *mux.Router) *mux.Router {
	router.Use(middleware.NoraGuard)

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

func getMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"memoryID": mux.Vars(req)["memoryID"],
	}

	if errs := validator.Nora.GetMemory(args); len(errs) > 0 {
		exception.Send(exception.Validation().AddErrors(errs...), writer)
		return
	}

	status, headers, body, err := business.Memory.Get(args)
	if err != nil {
		exception.Send(err, writer)
		return
	}
	if status == 0 {
		return
	}
	sendResponse(writer, status, headers, body)
}

func patchMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

func deleteMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"memoryID": mux.Vars(req)["memoryID"],
	}

	if errs := validator.Nora.DeleteMemory(args); len(errs) > 0 {
		exception.Send(exception.Validation().AddErrors(errs...), writer)
		return
	}

	status, headers, body, err := business.Memory.Delete(args)
	if err != nil {
		exception.Send(err, writer)
		return
	}
	if status == 0 {
		return
	}
	sendResponse(writer, status, headers, body)
}

func listMemoriesHandler(writer http.ResponseWriter, req *http.Request) {}

func postMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{}

	err := json.NewDecoder(req.Body).Decode(&args)
	if err != nil {
		exception.Send(exception.Validation().AddMessages("Invalid Body"), writer)
		return
	}

	if errs := validator.Nora.PostMemory(args); len(errs) > 0 {
		exception.Send(exception.Validation().AddErrors(errs...), writer)
		return
	}

	status, headers, body, err := business.Memory.Post(args)
	if err != nil {
		exception.Send(err, writer)
		return
	}
	if status == 0 {
		return
	}
	sendResponse(writer, status, headers, body)
}
