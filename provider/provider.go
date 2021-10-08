package provider

import (
	"context"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STATUSCAKE_APIKEY", nil),
				Description: "API Key for StatusCake",
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"statuscake_contact_group": ContactGroupResource(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (_ interface{}, dg diag.Diagnostics) {
	return statuscake.NewAPIClient(d.Get("apikey").(string)), dg
}
