package elasticache

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	versionStringRegexpInternalPattern = `[[:digit:]]+(.[[:digit:]]+){2}`
	versionStringRegexpPattern         = "^" + versionStringRegexpInternalPattern + "$"
)

var versionStringRegexp = regexp.MustCompile(versionStringRegexpPattern)

func validReplicationGroupAuthToken() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringLenBetween(16, 128),
		validation.StringMatch(regexp.MustCompile(`^[^@"\/]+$`), "must only contain alphanumeric characters or symbols (excl. @, \" and /)"),
	)
}

func validVersionString() schema.SchemaValidateFunc {
	return validation.StringMatch(versionStringRegexp, "must be a version string matching x.y.z")
}
