package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func validEngine() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"docdb",
	}, false)
}
