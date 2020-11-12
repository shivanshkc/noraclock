package validator

import validation "github.com/go-ozzo/ozzo-validation/v4"

type memoryRestValidator struct{}

// Memory : Struct that encapsulates all validation rules for Memory APIs.
var Memory = &memoryRestValidator{}

func (m *memoryRestValidator) GetMemoryByID(args map[string]interface{}) error {
	return validation.Validate(args, validation.Map(memoryIDParam))
}

func (m *memoryRestValidator) GetMemories(args map[string]interface{}) error {
	return validation.Validate(args, validation.Map(limitParam, offsetParam, skipBodyParam))
}

func (m *memoryRestValidator) PostMemory(args map[string]interface{}) error {
	return validation.Validate(args, validation.Map(titleParam, bodyParam))
}

func (m *memoryRestValidator) PatchMemory(args map[string]interface{}) error {
	if err := validation.Validate(args, validation.Map(memoryIDParam).AllowExtraKeys()); err != nil {
		return err
	}

	_, tExists := args["title"]
	_, bExists := args["body"]

	if !tExists && !bExists {
		return errEmptyUpdate
	}
	if tExists {
		if err := validation.Validate(args, validation.Map(titleParam).AllowExtraKeys()); err != nil {
			return err
		}
	}
	if bExists {
		if err := validation.Validate(args, validation.Map(bodyParam).AllowExtraKeys()); err != nil {
			return err
		}
	}

	return nil
}

func (m *memoryRestValidator) DeleteMemory(args map[string]interface{}) error {
	return validation.Validate(args, validation.Map(memoryIDParam))
}
