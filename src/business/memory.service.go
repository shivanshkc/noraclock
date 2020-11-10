package business

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
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
		return nil, err
	}

	return &Result{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func (m *memoryService) GetMemories(args map[string]interface{}) (*Result, error) {
	limit, offset := limitOffsetParser(args)

	memories, err := tables.Memory.Get(limit, offset)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]interface{}{"data": memories})
	if err != nil {
		return nil, err
	}

	return &Result{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func (m *memoryService) PostMemory(args map[string]interface{}) (*Result, error) {
	memoryID := uuid.New().String()

	err := tables.Memory.Insert(memoryID, args["title"].(string), args["body"].(string))
	if err != nil {
		return nil, err
	}

	return &Result{StatusCode: http.StatusCreated}, nil
}

func (m *memoryService) PatchMemory(args map[string]interface{}) (*Result, error) {
	err := tables.Memory.UpdateByID(args["memoryID"].(string), args)
	if err != nil {
		return nil, err
	}

	return &Result{StatusCode: http.StatusOK}, nil
}
