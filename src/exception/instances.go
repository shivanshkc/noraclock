package exception

import "net/http"

// BadRequest :
var BadRequest = func(message string) *Exception {
	return New(http.StatusBadRequest, "BAD_REQUEST", or(message, "bad request"))
}

// Validation : Validation Error for Client requests.
var Validation = func(message string) *Exception {
	return New(http.StatusBadRequest, "VALIDATION_ERROR", or(message, "validation error"))
}

// Unauthorized :
var Unauthorized = func(message string) *Exception {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", or(message, "unauthorized operation"))
}

// MemoryNotFound :
var MemoryNotFound = func(message string) *Exception {
	return New(http.StatusNotFound, "MEMORY_NOT_FOUND", or(message, "memory not found"))
}

// Unexpected : Shorthand for Internal Server Error
var Unexpected = func(message string) *Exception {
	return New(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", or(message, "unexpected error"))
}
