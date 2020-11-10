package business

import (
	"noraclock/v2/src/constants"
	"strconv"
)

func limitOffsetParser(args map[string]interface{}) (int64, int64) {
	limit, _ := args["limit"]
	offset, _ := args["offset"]

	strLimit, _ := limit.(string)
	strOffset, _ := offset.(string)

	var intLimit int64 = constants.MaxLimit
	var intOffset int64 = 0

	if strLimit != "" {
		intLimit, _ = strconv.ParseInt(strLimit, 10, 64)
	}
	if strOffset != "" {
		intOffset, _ = strconv.ParseInt(strOffset, 10, 64)
	}

	return intLimit, intOffset
}
