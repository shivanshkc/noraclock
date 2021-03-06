package exception

import "net/http"

// Validation : Thrown when invalid arguments received from client.
var Validation = func() *Exception {
	return &Exception{http.StatusBadRequest, "VALIDATION_ERROR", []string{}}
}

// Unauthorized : Thrown when an operation is performed without sufficient permissions.
var Unauthorized = func() *Exception {
	return &Exception{StatusCode: http.StatusUnauthorized, CustomCode: "UNAUTHORIZED_OPERATION_ERROR", Messages: []string{}}
}

// MemoryNotFound : Thrown when a required memory is not found.
var MemoryNotFound = func() *Exception {
	return &Exception{StatusCode: http.StatusNotFound, CustomCode: "MEMORY_NOT_FOUND", Messages: []string{}}
}

// Unexpected : Thrown when an unknown misbehaviour occurs.
var Unexpected = func() *Exception {
	return &Exception{http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", []string{}}
}
