package server

import "noraclock/src/logger"

var accLog = logger.Access()

type accessLogWriter struct{}

func (a *accessLogWriter) Write(input []byte) (int, error) {
	accLog.Sugar().Infow(string(input))
	return len(input), nil
}
