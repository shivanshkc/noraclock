package validator

import validation "github.com/go-ozzo/ozzo-validation/v4"

type memoryRestValidator struct{}

// Memory : Struct that encapsulates all validation rules for Memory APIs.
var Memory = &memoryRestValidator{}

func (m *memoryRestValidator) PostMemory(args map[string]interface{}) error {
	return validation.Validate(args, validation.Map(titleParam, bodyParam))
}
