package business

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"noraclock/src/logger"
	"noraclock/src/mongo"
	"strconv"
)

var log = logger.General()

// Memory implements the business methods of memory entity.
var Memory = &memory{}

type memory struct{}

func (m *memory) Get(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)

	memory, err := mongo.GetMemoryService().GetByID(memID)
	if err != nil {
		log.Sugar().Errorf("Memory.Get: Failed to get memory from database: %s", err.Error())
		return 0, nil, nil, err
	}

	memBytes, err := json.Marshal(memory)
	if err != nil {
		log.Sugar().Errorf("Memory.Get: Failed to marshal memory: %s", err.Error())
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, []byte(fmt.Sprintf(`{"data":%s}`, string(memBytes))), nil
}

func (m *memory) Post(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := uuid.New().String()

	memory := &mongo.Memory{
		ID:           memID,
		Title:        args["title"].(string),
		Body:         args["body"].(string),
	}

	err := mongo.GetMemoryService().Create(memory)
	if err != nil {
		log.Sugar().Errorf("Memory.Post: Failed to create memory: %s", err.Error())
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, []byte(fmt.Sprintf(`{"id":"%s"}`, memID)), nil
}

func (m *memory) Delete(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)

	err := mongo.GetMemoryService().Delete(memID)
	if err != nil {
		log.Sugar().Errorf("Memory.Delete: Failed to delete memory from database: %s", err.Error())
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, nil, nil
}

func (m *memory) Patch(args map[string]interface{}) (int, map[string]string, []byte, error) {
	memID := args["memoryID"].(string)
	title, tExists := args["title"]
	body, bExists := args["body"]

	memory := &mongo.Memory{ID: memID}
	if tExists {
		memory.Title = title.(string)
	}
	if bExists {
		memory.Body = body.(string)
	}

	err := mongo.GetMemoryService().Update(memory)
	if err != nil {
		log.Sugar().Errorf("Memory.Patch : Failed to update memory in database: %s", err.Error())
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, nil, nil
}

func (m *memory) List(args map[string]interface{}) (int, map[string]string, []byte, error) {
	limitIn, lExists := args["limit"]
	offsetIn, oExists := args["offset"]
	skipBodyIn, sExists := args["skipBody"]

	var limit, offset int64
	var skipBody bool

	if lExists && limitIn != "" {
		limit, _ = strconv.ParseInt(limitIn.(string), 10, 64)
	} else {
		limit = 100
	}
	if oExists && offsetIn != "" {
		offset, _ = strconv.ParseInt(offsetIn.(string), 10, 64)
	} else {
		offset = 0
	}
	if sExists && skipBodyIn != "" {
		skipBody, _ = strconv.ParseBool(skipBodyIn.(string))
	} else {
		skipBody = false
	}

	memories, err := mongo.GetMemoryService().GetList(limit, offset)
	if err != nil {
		log.Sugar().Errorf("Memory.List : Failed to list memories in database: %s", err.Error())
		return 0, nil, nil, err
	}

	count, err := mongo.GetMemoryService().GetCount()
	if err != nil {
		log.Sugar().Errorf("Memory.List : Failed to count memories in database: %s", err.Error())
		return 0, nil, nil, err
	}

	if skipBody {
		for _, mem := range memories {
			mem.Body = ""
		}
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
