package business

import (
	"net/http"
	"noraclock/v2/src/logger"
)

var log, _ = logger.Get()

// Result : Return type of every business method.
type Result struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

// Send : Accepts a ResponseWriter and writes itself using it.
func (hr *Result) Send(writer http.ResponseWriter) {
	writer.WriteHeader(hr.StatusCode)

	for key, value := range hr.Headers {
		writer.Header().Set(key, value)
	}

	_, err := writer.Write(hr.Body)
	if err != nil {
		log.Sugar().Errorf("Failed to write response. %s", err.Error())
	}
}
