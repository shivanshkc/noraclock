package exception

import (
	"encoding/json"
	"net/http"
	"noraclock/src/logger"
)

var log = logger.General()

// Exception : The custom error type, implements the error interface.
type Exception struct {
	StatusCode int      `json:"status_code"`
	CustomCode string   `json:"custom_code"`
	Messages   []string `json:"messages"`
}

func (e *Exception) Error() string {
	return e.CustomCode
}

// AddMessages : Chainable method to add string messages to the Exception.
func (e *Exception) AddMessages(message ...string) *Exception {
	e.Messages = append(e.Messages, message...)
	return e
}

// AddErrors : Chainable method to add error messages to the Exception.
func (e *Exception) AddErrors(errors ...error) *Exception {
	for _, err := range errors {
		e.Messages = append(e.Messages, err.Error())
	}
	return e
}

// ToJSON : Converts the Exception to JSON.
func (e *Exception) ToJSON() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		log.Sugar().Warnf("exception.ToJSON: Failed to marshal Exception to JSON: %s", err.Error())
		return nil
	}
	return bytes
}

// Send : Sends an error safely as an HTTP response.
func Send(err error, writer http.ResponseWriter) {
	exc, ok := err.(*Exception)
	if !ok {
		log.Sugar().Errorf("exception.Send: Unexpected error: %s", err.Error())
		exc = Unexpected()
	}

	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(exc.StatusCode)
	_, wErr := writer.Write(exc.ToJSON())
	if wErr != nil {
		log.Sugar().Warnf("exception.Send: Failed to write response: %s", err.Error())
	}
}
