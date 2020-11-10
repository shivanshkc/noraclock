package router

import (
	"net/http"
	"noraclock/v2/src/business"
	"noraclock/v2/src/exception"
	"noraclock/v2/src/validator"
)

func getMemoryByIDHandler(writer http.ResponseWriter, req *http.Request) {}

func getMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

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

func patchMemoryHandler(writer http.ResponseWriter, req *http.Request) {}

func deleteMemoryHandler(writer http.ResponseWriter, req *http.Request) {}
