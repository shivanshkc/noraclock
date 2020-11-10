package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"noraclock/v2/src/business"
	"noraclock/v2/src/exception"
	"noraclock/v2/src/validator"
)

func getMemoryByIDHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"memoryID": mux.Vars(req)["memoryID"],
	}

	if err := validator.Memory.GetMemoryByID(args); err != nil {
		sendError(writer, exception.Validation(err.Error()))
		return
	}

	result, err := business.Memory.GetMemoryByID(args)
	if err != nil {
		sendError(writer, err)
		return
	}
	result.Send(writer)
}

func getMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"limit":  req.URL.Query().Get("limit"),
		"offset": req.URL.Query().Get("offset"),
	}

	if err := validator.Memory.GetMemories(args); err != nil {
		sendError(writer, exception.Validation(err.Error()))
		return
	}

	result, err := business.Memory.GetMemories(args)
	if err != nil {
		sendError(writer, err)
		return
	}
	result.Send(writer)
}

func postMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args, err := readBody(req)
	if err != nil {
		sendError(writer, exception.BadRequest(err.Error()))
		return
	}

	if err = validator.Memory.PostMemory(args); err != nil {
		sendError(writer, exception.Validation(err.Error()))
		return
	}

	result, err := business.Memory.PostMemory(args)
	if err != nil {
		sendError(writer, err)
		return
	}
	result.Send(writer)
}

func patchMemoryHandler(writer http.ResponseWriter, req *http.Request) {
	args := map[string]interface{}{
		"memoryID": mux.Vars(req)["memoryID"],
	}

	body, err := readBody(req)
	if err != nil {
		sendError(writer, exception.Validation(err.Error()))
		return
	}

	for key, value := range body {
		args[key] = value
	}

	if err = validator.Memory.PatchMemory(args); err != nil {
		sendError(writer, exception.Validation(err.Error()))
		return
	}
}

func deleteMemoryHandler(writer http.ResponseWriter, req *http.Request) {}
