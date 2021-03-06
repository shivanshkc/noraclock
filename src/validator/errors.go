package validator

import (
	"fmt"
	"noraclock/src/constants"
)

var (
	errMemoryID = fmt.Errorf("'memoryID' should follow:: UUIDv4 string")
	errTitle    = fmt.Errorf("'title' should follow:: max-length: %d & min-length: %d", constants.MemTitleMaxLen, constants.MemTitleMinLen)
	errBody     = fmt.Errorf("'body' should follow:: max-length: %d & min-length: %d", constants.MemBodyMaxLen, constants.MemBodyMinLen)
)
