package business

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"noraclock/src/configs"
	"noraclock/src/constants"
	"noraclock/src/database"
	"noraclock/src/exception"
	"noraclock/src/logger"
	"time"
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
	memID := uuid.New().String()

	docMap := map[string]interface{}{
		"id":        memID,
		"title":     args["title"],
		"body":      args["body"],
		"createdAt": time.Now().Unix() * 1000,
		"updatedAt": time.Now().Unix() * 1000,
	}

	docBytes, err := json.Marshal(docMap)
	if err != nil {
		return 0, nil, nil, err
	}

	err = database.CouchDB.CreateDoc(conf.CouchDB.Database, memID, docBytes)
	if err == nil {
		return http.StatusCreated, nil, []byte(fmt.Sprintf(`{"id":"%s"}`, memID)), nil
	}
	if err.Error() == constants.CouchUpdateConflictReason {
		return 0, nil, nil, exception.MemoryAlreadyExists()
	}
	return 0, nil, nil, err
}
