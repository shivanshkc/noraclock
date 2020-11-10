package business

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"noraclock/v2/src/exception"
	"noraclock/v2/src/tables"
)

type memoryService struct{}

// Memory : Struct that encapsulates all Memory business methods.
var Memory = &memoryService{}

func (m *memoryService) GetMemoryByID(args map[string]interface{}) (*Result, error) {
	memory, err := tables.Memory.GetByID(args["memoryID"].(string))
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]interface{}{"data": memory})
	if err != nil {
		return nil, exception.Unexpected(err.Error())
	}

	return &Result{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func (m *memoryService) GetMemories(args map[string]interface{}) (*Result, error) {
	return nil, nil
}

func (m *memoryService) PostMemory(args map[string]interface{}) (*Result, error) {
	memoryID := uuid.New().String()

	err := tables.Memory.Insert(memoryID, args["title"].(string), args["body"].(string))
	if err != nil {
		return nil, err
	}

	return &Result{StatusCode: http.StatusCreated}, nil
}
