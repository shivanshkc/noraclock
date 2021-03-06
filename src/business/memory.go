package business

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
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
	if err != nil {
		if err.Error() == constants.CouchMissingReason || err.Error() == constants.CouchDeletedReason {
			return 0, nil, nil, exception.MemoryNotFound()
		}
		log.Sugar().Errorf("Memory.Get: Failed to get memory from database: %s", err.Error())
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, []byte(fmt.Sprintf(`{"data":%s}`, string(doc))), nil
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

func (m *memory) Delete(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)

	err := database.CouchDB.DeleteDoc(conf.CouchDB.Database, memID)
	if err == nil {
		return http.StatusOK, nil, nil, nil
	}
	if err.Error() == constants.CouchMissingReason || err.Error() == constants.CouchDeletedReason {
		return 0, nil, nil, exception.MemoryNotFound()
	}
	log.Sugar().Errorf("Memory.Delete: Failed to delete memory from database: %s", err.Error())
	return 0, nil, nil, err
}

func (m *memory) Patch(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)
	title, tExists := args["title"]
	body, bExists := args["body"]

	doc, err := database.CouchDB.GetDoc(conf.CouchDB.Database, memID)
	if err != nil {
		if err.Error() == constants.CouchMissingReason || err.Error() == constants.CouchDeletedReason {
			return 0, nil, nil, exception.MemoryNotFound()
		}
		log.Sugar().Errorf("Memory.Patch : Failed to retrieve current memory from database: %s", err.Error())
		return 0, nil, nil, err
	}

	docMap := map[string]interface{}{}
	if err := json.Unmarshal(doc, &docMap); err != nil {
		log.Sugar().Errorf("Memory.Patch : Failed to unmarshal memory doc from database: %s", err.Error())
		return 0, nil, nil, err
	}

	docMap["updatedAt"] = time.Now().Unix() * 1000
	if tExists {
		docMap["title"] = title
	}
	if bExists {
		docMap["body"] = body
	}

	docBytes, err := json.Marshal(docMap)
	if err != nil {
		log.Sugar().Errorf("Memory.Patch : Failed to marshal updated memory: %s", err.Error())
		return 0, nil, nil, err
	}
	err = database.CouchDB.UpdateDocWithRev(conf.CouchDB.Database, memID, docMap["_rev"].(string), docBytes)
	if err != nil {
		log.Sugar().Errorf("Memory.Patch : Failed to update memory in database: %s", err.Error())
		return 0, nil, nil, err
	}
	return http.StatusOK, nil, nil, nil
}

func (m *memory) List(args map[string]interface{}) (int, map[string]string, []byte, error) {
	limit, lExists := args["limit"]
	offset, oExists := args["offset"]
	skipBody, sExists := args["skipBody"]

	if !lExists || limit == "" {
		limit = "100"
	}
	if !oExists || offset == "" {
		offset = "0"
	}
	if !sExists || skipBody == "" {
		skipBody = false
	}

	query := url.Values{}
	query.Set("limit", limit.(string))
	query.Set("skip", offset.(string))
	query.Set("include_docs", "true")

	_, count, results, err := database.CouchDB.GetDocsByView(conf.CouchDB.Database, constants.CouchDesign, constants.CouchListMemoriesView, query)
	if err != nil {
		return 0, nil, nil, err
	}

	var memories []interface{}
	for _, result := range results {
		resultMap, ok := result.(map[string]interface{})
		if !ok {
			continue
		}
		doc, exists := resultMap["doc"]
		if !exists {
			continue
		}
		docMap, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		if skipBody == "false" {
			delete(docMap, "body")
		}
		memories = append(memories, docMap)
	}

	body, err := json.Marshal(map[string]map[string]interface{}{
		"data": {"count": count, "docs": memories},
	})
	if err != nil {
		log.Sugar().Errorf("Memory.List: Failed to marshal response body: %s", err.Error())
		return 0, nil, nil, err
	}
	return http.StatusOK, nil, body, nil
}
