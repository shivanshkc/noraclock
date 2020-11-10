package server

import (
	"net/http"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/logger"
	"noraclock/v2/src/router"
)

// Start : Initiates the HTTP Server and calls ListenAndServe on it.
func Start() error {
	log, _ := logger.Get()

	log.Sugar().Infof(
		"Starting %s:%s HTTP server at %s:%s",
		configs.Service.Name,
		configs.Service.Version,
		configs.Server.Address,
		configs.Server.Port,
	)

	return http.ListenAndServe(configs.Server.Address+":"+configs.Server.Port, router.Handlers())
}
