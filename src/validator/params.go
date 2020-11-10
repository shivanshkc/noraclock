package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"noraclock/v2/src/constants"
	"regexp"
	"strconv"
)

var memoryIDParam = validation.Key(
	"memoryID",
	validation.Required,
	is.UUIDv4,
)

var titleParam = validation.Key(
	"title",
	validation.Required,
	validation.Match(regexp.MustCompile(constants.MatchAllRegex)),
	validation.Length(constants.TitleMinLength, constants.TitleMaxLength),
)

var bodyParam = validation.Key(
	"body",
	validation.Match(regexp.MustCompile(constants.MatchAllRegex)),
	validation.Length(constants.BodyMinLength, constants.BodyMaxLength),
)

var limitParam = validation.Key(
	"limit",
	validation.By(func(value interface{}) error {
		strLimit, ok := value.(string)
		if !ok {
			return errLimitType
		}

		if strLimit == "" {
			return nil
		}

		intLimit, err := strconv.ParseInt(strLimit, 10, 64)
		if err != nil {
			return errLimitType
		}

		if intLimit < 1 || intLimit > constants.MaxLimit {
			return errLimitRange
		}
		return nil
	}),
)

var offsetParam = validation.Key(
	"offset",
	validation.By(func(value interface{}) error {
		strOffset, ok := value.(string)
		if !ok {
			return errOffset
		}

		if strOffset == "" {
			return nil
		}

		intOffset, err := strconv.ParseInt(strOffset, 10, 64)
		if err != nil {
			return errOffset
		}

		if intOffset < 0 {
			return errOffset
		}
		return nil
	}),
)
