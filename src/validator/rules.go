package validator

import (
	"github.com/shivanshkc/valkyrie/v2"
	"noraclock/src/constants"
)

var (
	memoryIDRule = valkyrie.PureString().UUIDv4().WithError(errMemoryID)
	titleRule    = valkyrie.PureString().LenLTE(constants.MemTitleMaxLen).LenGTE(constants.MemTitleMinLen).WithError(errTitle)
	bodyRule     = valkyrie.PureString().LenLTE(constants.MemBodyMaxLen).LenGTE(constants.MemBodyMinLen).WithError(errBody)
	limitRule    = valkyrie.StringInt().Allow("").GTE(1).LTE(100).WithError(errLimit)
	offsetRule   = valkyrie.StringInt().Allow("").GTE(0).WithError(errOffset)
	skipBodyRule = valkyrie.PureString().Blind().Allow("", "true", "false").WithError(errSkipBody)
)
