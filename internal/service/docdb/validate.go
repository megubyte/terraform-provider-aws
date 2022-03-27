package docdb

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func validClusterIdentifier() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
	)
}

func validClusterSnapshotIdentifier() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
	)
}

func validEngine() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"docdb",
	}, false)
}

func validIdentifier() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
		validation.StringLenBetween(1, 63),
	)
}

func validIdentifierPrefix() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
	)
}

func validParamGroupName() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringDoesNotMatch(regexp.MustCompile(`-$`), "must not end with a hyphen"),
		validation.StringLenBetween(1, 255),
	)
}

func validGlobalCusterIdentifier() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringLenBetween(1, 255),
	)
}

func validParamGroupNamePrefix() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`[0-9a-z-]+$`), "must contain only lowercase alphanumeric characters and hyphens"),
		validation.StringMatch(regexp.MustCompile(`^[a-z]`), "must start a lowercase letter"),
		validation.StringDoesNotMatch(regexp.MustCompile(`--`), "must not contain two consecutive hyphens"),
		validation.StringLenBetween(1, 255),
	)
}

func validSubnetGroupName() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`^[ .0-9a-z-_]+$`), "must contain only alphanumeric characters, hyphens, underscores, spaces and periods"),
		validation.StringDoesNotMatch(regexp.MustCompile(`^default$`), "must not be 'default'"),
		validation.StringLenBetween(1, 255),
	)
}

func validSubnetGroupNamePrefix() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringMatch(regexp.MustCompile(`^[ .0-9a-z-_]+$`), "must contain only alphanumeric characters, hyphens, underscores, spaces and periods"),
		validation.StringLenBetween(1, 255-resource.UniqueIDSuffixLength),
	)
}
