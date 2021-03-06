package validator

import (
	"fmt"
	"noraclock/src/constants"
)

var (
	errEmptyUpdate = fmt.Errorf("no update parameters provided")
	errMemoryID    = fmt.Errorf("'memoryID' should follow:: UUIDv4 string")
	errTitle       = fmt.Errorf("'title' should follow:: max-length: %d & min-length: %d", constants.MemTitleMaxLen, constants.MemTitleMinLen)
	errBody        = fmt.Errorf("'body' should follow:: max-length: %d & min-length: %d", constants.MemBodyMaxLen, constants.MemBodyMinLen)
	errLimit       = fmt.Errorf("'limit' should follow:: integer & max: 100 & min: 1")
	errOffset      = fmt.Errorf("'offset' should follow:: integer & min: 0")
	errSkipBody    = fmt.Errorf("'skipBody' should follow:: boolean")
)
