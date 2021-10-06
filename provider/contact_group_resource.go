package provider

import (
	"context"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ContactGroupResource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   updateContactGroup,
		CreateContext: createContactGroup,
		UpdateContext: updateContactGroup,
		DeleteContext: deleteContactGroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ping_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mobile_numbers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"email_addresses": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"integration_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func readContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	res, err := client.GetContactGroup(ctx, d.Id()).Execute()
	if err != nil {
		// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return diag.FromErr(err)
	}

	d.Set("name", res.Data.Name)
	d.Set("mobile_numbers", res.Data.MobileNumbers)
	d.Set("email_addresses", res.Data.EmailAddresses)
	d.Set("integration_ids", res.Data.Integrations)
	d.Set("ping_url", res.Data.PingURL)

	return dg
}

func createContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	res, err := client.CreateContactGroup(ctx).
		Name(d.Get("name").(string)).
		PingURL(d.Get("ping_url").(string)).
		MobileNumbers(castSetToSliceStrings(d.Get("mobile_numbers").(*schema.Set).List())).
		EmailAddresses(castSetToSliceStrings(d.Get("email_addresses").(*schema.Set).List())).
		Integrations(castSetToSliceStrings(d.Get("integration_ids").(*schema.Set).List())).
		Execute()

	if err != nil {
		// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return diag.FromErr(err)
	}

	d.SetId(res.Data.NewID)

	return readContactGroup(ctx, d, i)
}

func updateContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	err := client.UpdateContactGroup(ctx, d.Id()).
		Name(d.Get("name").(string)).
		PingURL(d.Get("ping_url").(string)).
		MobileNumbers(castSetToSliceStrings(d.Get("mobile_numbers").(*schema.Set).List())).
		EmailAddresses(castSetToSliceStrings(d.Get("email_addresses").(*schema.Set).List())).
		Integrations(castSetToSliceStrings(d.Get("integration_ids").(*schema.Set).List())).
		Execute()

	if err != nil {
		// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return diag.FromErr(err)
	}

	return readContactGroup(ctx, d, i)
}

func deleteContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	err := client.DeleteContactGroup(ctx, d.Id()).Execute()
	if err != nil {
		// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return diag.FromErr(err)
	}

	return dg
}

func orEmptySlice(a []string) []string {
	if a == nil || len(a) == 0 {
		return []string{}
	}

	return a
}

func castSetToSliceStrings(configured []interface{}) []string {
	res := make([]string, len(configured))

	for i, element := range configured {
		res[i] = element.(string)
	}
	return res
}
