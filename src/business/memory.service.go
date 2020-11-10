package business

import (
	"github.com/google/uuid"
	"net/http"
	"noraclock/v2/src/tables"
)

type memoryService struct{}

// Memory : Struct that encapsulates all Memory business methods.
var Memory = &memoryService{}

func (m *memoryService) PostMemory(args map[string]interface{}) (*Result, error) {
	memoryID := uuid.New().String()

	err := tables.Memory.Insert(memoryID, args["title"].(string), args["body"].(string))
	if err != nil {
		return nil, err
	}

	return &Result{StatusCode: http.StatusCreated}, nil
}
