package router

import (
	"encoding/json"
	"net/http"
)

func sendResponse(writer http.ResponseWriter, statusCode int, headers map[string]string, body interface{}) {
	if statusCode == 0 {
		log.Sugar().Infof("router.sendResponse: Status Code is 0. Response is assumed to be already sent.")
		return
	}

	contentType := getContentType(body)
	writer.Header().Set("content-type", contentType)

	if headers != nil {
		for key, value := range headers {
			writer.Header().Set(key, value)
		}
	}

	writer.WriteHeader(statusCode)

	bodyBytes, err := body2Bytes(body)
	if err != nil {
		log.Sugar().Errorf("router.sendResponse: Failed to convert response body to bytes: %s", err.Error())
		return
	}

	if _, err := writer.Write(bodyBytes); err != nil {
		log.Sugar().Errorf("router.sendResponse: Failed to write response body: %s", err.Error())
		return
	}
}

func getContentType(body interface{}) string {
	switch body.(type) {
	case string:
		return "text/html"
	default:
		return "application/json"
	}
}

func body2Bytes(body interface{}) ([]byte, error) {
	switch body.(type) {
	case string:
		return []byte(body.(string)), nil
	case []byte:
		return body.([]byte), nil
	default:
		return json.Marshal(body)
	}
}
