package configs

import "github.com/hisitra/confine/v2"

var configMap = map[string]interface{}{
	"src/configs/json/logger.json":   Logger,
	"src/configs/json/postgres.json": Postgres,
	"src/configs/json/server.json":   Server,
	"src/configs/json/service.json":  Service,
}

func init() {
	err := confine.LoadMany(configMap)
	if err != nil {
		panic(err)
	}
}
