package provider

import (
	"context"
	"fmt"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ContactGroupResource - Schema
func ContactGroupResource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   readContactGroup,
		CreateContext: createContactGroup,
		UpdateContext: updateContactGroup,
		DeleteContext: deleteContactGroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"ping_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
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
		return Prettify(dg, fmt.Sprintf("Error from Statuscake API when reading contact group %s", d.Id()), err, true)
	}

	if err := d.Set("name", res.Data.Name); err != nil {
		return Prettify(dg, fmt.Sprintf("Error setting property name for %s", d.Id()), err, false)
	}
	if err := d.Set("ping_url", res.Data.PingURL); err != nil {
		return Prettify(dg, fmt.Sprintf("Error setting property name for %s", d.Id()), err, false)
	}
	if err := d.Set("mobile_numbers", res.Data.MobileNumbers); err != nil {
		return Prettify(dg, fmt.Sprintf("Error setting property name for %s", d.Id()), err, false)
	}
	if err := d.Set("email_addresses", res.Data.EmailAddresses); err != nil {
		return Prettify(dg, fmt.Sprintf("Error setting property name for %s", d.Id()), err, false)
	}
	if err := d.Set("integration_ids", res.Data.Integrations); err != nil {
		return Prettify(dg, fmt.Sprintf("Error setting property name for %s", d.Id()), err, false)
	}

	return dg
}

func createContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	res, err := client.CreateContactGroup(ctx).
		Name(d.Get("name").(string)).
		PingURL(d.Get("ping_url").(string)).
		MobileNumbers(normalize(d.Get("mobile_numbers").(*schema.Set).List())).
		EmailAddresses(normalize(d.Get("email_addresses").(*schema.Set).List())).
		Integrations(normalize(d.Get("integration_ids").(*schema.Set).List())).
		Execute()
	if err != nil {
		return Prettify(dg, "Error from Statuscake API when creating contact group", err, true)
	}

	d.SetId(res.Data.NewID)

	return readContactGroup(ctx, d, i)
}

func updateContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	// Only run update in case there are updated fields
	if d.HasChanges("name", "ping_url", "mobile_numbers", "email_addresses", "integration_ids") {
		err := client.UpdateContactGroup(ctx, d.Id()).
			Name(d.Get("name").(string)).
			PingURL(d.Get("ping_url").(string)).
			MobileNumbers(normalize(d.Get("mobile_numbers").(*schema.Set).List())).
			EmailAddresses(normalize(d.Get("email_addresses").(*schema.Set).List())).
			Integrations(normalize(d.Get("integration_ids").(*schema.Set).List())).
			Execute()

		if err != nil {
			return Prettify(dg, fmt.Sprintf("Error from Statuscake API when updating contact group %s", d.Id()), err, true)
		}
	}

	return readContactGroup(ctx, d, i)
}

func deleteContactGroup(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*statuscake.APIClient)

	err := client.DeleteContactGroup(ctx, d.Id()).Execute()
	if err != nil {
		return Prettify(dg, fmt.Sprintf("Error from Statuscake API when deleting contact group %s", d.Id()), err, true)
	}

	return dg
}

// Normalize list of strings
func normalize(i []interface{}) []string {
	r := make([]string, len(i))
	for e, el := range i {
		r[e] = el.(string)
	}
	return r
}
