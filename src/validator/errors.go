package validator

import (
	"errors"
	"fmt"
	"noraclock/v2/src/constants"
)

var errLimitType = errors.New("limit should be an integer")
var errLimitRange = fmt.Errorf("limit should be between 1 and %d", constants.MaxLimit)

var errOffset = errors.New("offset should be a non-negative integer")

var errEmptyUpdate = errors.New("no update parameters provided")
