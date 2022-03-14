package eks

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func validClusterName() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringLenBetween(1, 100),
		validation.StringMatch(regexp.MustCompile(`^[0-9A-Za-z][A-Za-z0-9\-_]+$`), "must contain alphanumeric characters (dashes and underscores allowed other than at the start)"),
	)
}
