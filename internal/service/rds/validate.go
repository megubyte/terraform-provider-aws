package rds

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validEventSubscriptionName = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[0-9A-Za-z-]+$`), "must only contain alphanumeric characters and hyphens"),
	validation.StringLenBetween(1, 255),
)

var validOptionGroupName = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumeric characters and hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
	validation.StringLenBetween(1, 255),
)

var validOptionGroupNamePrefix = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumeric characters and hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	validation.StringLenBetween(1, 229),
)

var validParamGroupName = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumeric characters and hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
	validation.StringLenBetween(1, 255),
)

var validParamGroupNamePrefix = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumeric characters and hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	validation.StringLenBetween(1, 255),
)

var validSubnetGroupName = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[ .0-9a-z-_]+$`), "must only contain lowercase alphanumeric characters, hyphens, underscores, periods, and spaces"),
	validation.StringLenBetween(1, 255),
	validation.StringMatch(regexp.MustCompile(`(?i)^default$`), "is not allowed as Default"),
)

var validSubnetGroupNamePrefix = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[ .0-9a-z-_]+$`), "must contain only lowercase alphanumeric characters, hyphens, underscores, periods, and spaces"),
	validation.StringLenBetween(1, 229),
)

var validEngine = validation.StringInSlice([]string{
	"aurora",
	"aurora-mysql",
	"aurora-postgresql",
}, false)

var validIdentifier = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumeric characters and hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
)

var validIdentifierPrefix = validation.All(
	validation.StringMatch(regexp.MustCompile(`^[0-9a-z-]+$`), "must only contain lowercase alphanumetic characters and hyphens"),
	validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must have a first character which is a letter"),
	validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
)
