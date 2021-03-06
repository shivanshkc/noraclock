package server

import (
	"github.com/gorilla/handlers"
	"net/http"
	"noraclock/src/configs"
	"noraclock/src/logger"
	"noraclock/src/router"
)

var conf = configs.Get()
var log = logger.General()

// Start initiates the HTTP server.
func Start() {
	log.Sugar().Infof(
		"server.Start: Starting %s@%s HTTP server at %s",
		conf.Service.Name,
		conf.Service.Version,
		conf.Server.Address,
	)

	err := http.ListenAndServe(conf.Server.Address, handlers.CombinedLoggingHandler(&accessLogWriter{}, router.Get()))
	if err != nil {
		panic(err)
	}
}
