package business

import (
	"net/http"
	"noraclock/src/configs"
	"noraclock/src/constants"
	"noraclock/src/database"
	"noraclock/src/exception"
	"noraclock/src/logger"
)

var conf = configs.Get()
var log = logger.General()

// Memory implements the business methods of memory entity.
var Memory = &memory{}

type memory struct{}

func (m *memory) Get(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)

	doc, err := database.CouchDB.GetDoc(conf.CouchDB.Database, memID)
	if err == nil {
		return http.StatusOK, nil, doc, nil
	}
	if err.Error() == constants.CouchMissingReason {
		return 0, nil, nil, exception.MemoryNotFound()
	}
	log.Sugar().Errorf("Memory.Get: Failed to get memory from database: %s", err.Error())
	return 0, nil, nil, err
}

func (m *memory) Post(args map[string]interface{}) (int, map[string]string, []byte, error) {
	return http.StatusOK, nil, nil, nil
}
