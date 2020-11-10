package router

import (
	"encoding/json"
	"net/http"
	exc "noraclock/v2/src/exception"
	"noraclock/v2/src/logger"
)

func sendError(writer http.ResponseWriter, err error) {
	log, _ := logger.Get()

	exception, ok := err.(*exc.Exception)
	if !ok {
		log.Sugar().Warn("Sending non-exception error in response.")
		exception = exc.Unexpected(err.Error())
	}

	exception.Send(writer)
}

func readBody(req *http.Request) (map[string]interface{}, error) {
	body := map[string]interface{}{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
