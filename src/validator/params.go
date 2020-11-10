package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"noraclock/v2/src/constants"
	"regexp"
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
	validation.Length(constants.TitleMinLength, constants.TitleMaxLength),
	validation.Length(constants.BodyMinLength, constants.BodyMaxLength),
)
