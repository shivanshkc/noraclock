package validator

// Nora implements the validations methods for NoraAccess APIs.
var Nora = &noraValidator{}

type noraValidator struct{}

func (n *noraValidator) GetMemory(args map[string]interface{}) []error {
	var errs []error

	memID, exists := args["memoryID"]
	if !exists {
		errs = append(errs, errMemoryID)
	} else if err := memoryIDRule.Apply(memID); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (n *noraValidator) PostMemory(args map[string]interface{}) []error {
	var errs []error

	title, exists := args["title"]
	if !exists {
		errs = append(errs, errTitle)
	} else if err := titleRule.Apply(title); err != nil {
		errs = append(errs, err)
	}

	body, exists := args["body"]
	if !exists {
		errs = append(errs, errBody)
	} else if err := bodyRule.Apply(body); err != nil {
		errs = append(errs, err)
	}

	return errs
}
