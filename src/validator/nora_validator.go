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

func (n *noraValidator) DeleteMemory(args map[string]interface{}) []error {
	var errs []error

	memID, exists := args["memoryID"]
	if !exists {
		errs = append(errs, errMemoryID)
	} else if err := memoryIDRule.Apply(memID); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (n *noraValidator) PatchMemory(args map[string]interface{}) []error {
	var errs []error

	memID, exists := args["memoryID"]
	if !exists {
		errs = append(errs, errMemoryID)
	} else if err := memoryIDRule.Apply(memID); err != nil {
		errs = append(errs, err)
	}

	title, tExists := args["title"]
	if tExists {
		if err := titleRule.Apply(title); err != nil {
			errs = append(errs, err)
		}
	}

	body, bExists := args["body"]
	if bExists {
		if err := bodyRule.Apply(body); err != nil {
			errs = append(errs, err)
		}
	}

	if !tExists && !bExists {
		errs = append(errs, errEmptyUpdate)
	}
	return errs
}

func (n *noraValidator) ListMemories(args map[string]interface{}) []error {
	var errs []error

	return errs
}
