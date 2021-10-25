package provider

import (
	"fmt"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Prettify - Inject statuscake/terraform context into errors
func Prettify(dgs diag.Diagnostics, m string, e error, hydrate bool) diag.Diagnostics {
	d := e.Error()
	if hydrate {
		d = fmt.Sprintf(
			"Code: %d, Message: %s, Errors: %s",
			e.(statuscake.APIError).Status,
			e.(statuscake.APIError).Message,
			e.(statuscake.APIError).Errors,
		)
	}
	return append(dgs, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  m,
		Detail:   d,
	})
}
