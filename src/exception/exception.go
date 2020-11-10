package exception

import (
	"encoding/json"
	"fmt"
	"net/http"
	"noraclock/v2/src/logger"
)

var log, _ = logger.Get()

// Exception : Exception type implements the error interface.
type Exception struct {
	Code       int    `json:"code"`
	CustomCode string `json:"customCode"`
	Message    string `json:"message"`
}

func (e *Exception) Error() string {
	return e.Message
}

// ToJSON : Converts the Exception type to JSON byte array.
func (e *Exception) ToJSON() []byte {
	bytes, err := json.Marshal(e)
	if err == nil {
		return bytes
	}

	jsonStr := fmt.Sprintf(
		`{"code":%d,"customCode":"%s","message":"%s"}`,
		e.Code,
		e.CustomCode,
		e.Message,
	)

	return []byte(jsonStr)
}

// Send : Accepts a ResponseWriter and sends the Exception as Response using it.
func (e *Exception) Send(writer http.ResponseWriter) {
	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(e.Code)
	_, err := writer.Write(e.ToJSON())
	if err != nil {
		log.Sugar().Errorf("Failed to write response. %s", err.Error())
	}
}

// New : To create a new Exception.
func New(code int, customCode string, message string) *Exception {
	return &Exception{code, customCode, message}
}
