package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"noraclock/v2/src/constants"
	"regexp"
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
