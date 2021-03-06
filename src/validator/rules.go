package validator

import (
	"github.com/shivanshkc/valkyrie/v2"
	"noraclock/src/constants"
)

var (
	memoryIDRule = valkyrie.PureString().UUIDv4().WithError(errMemoryID)
	titleRule    = valkyrie.PureString().LenLTE(constants.MemTitleMaxLen).LenGTE(constants.MemTitleMaxLen)
	bodyRule     = valkyrie.PureString().LenLTE(constants.MemBodyMaxLen).LenGTE(constants.MemBodyMaxLen)
)
