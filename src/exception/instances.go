package exception

import "net/http"

// UnknownLogLevel : Thrown while initiating the logger.
var UnknownLogLevel = func(message string) *Exception {
	return New(http.StatusBadRequest, "UNKNOWN_LOG_LEVEL", or(message, "unknown log level"))
}

// ProvidedPathIsDir : Thrown while initiating the logger.
var ProvidedPathIsDir = func(message string) *Exception {
	return New(http.StatusBadRequest, "PROVIDED_PATH_IS_DIR", or(message, "provided path is a directory"))
}

// Validation : Validation Error for Client requests.
var Validation = func(message string) *Exception {
	return New(http.StatusBadRequest, "VALIDATION_ERROR", or(message, "validation error"))
}

// Unauthorized :
var Unauthorized = func(message string) *Exception {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", or(message, "unauthorized operation"))
}

// Unexpected : Shorthand for Internal Server Error
var Unexpected = func(message string) *Exception {
	return New(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", or(message, "unexpected error"))
}
