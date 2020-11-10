package business

type memoryService struct{}

// Memory : Struct that encapsulates all Memory business methods.
var Memory = &memoryService{}

func (m *memoryService) PostMemory(args map[string]interface{}) (*Result, error) {
	return &Result{
		StatusCode: 200,
		Headers:    nil,
		Body:       nil,
	}, nil
}
