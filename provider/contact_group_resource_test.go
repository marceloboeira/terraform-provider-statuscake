package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccCheckContactGroupDestroy(s *terraform.State) error {
	res, err := APIClient().ListContactGroups(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("Cannot list contact groups: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "statuscake_contact_group" {
			continue
		}

		if len(res.Data) > 0 {
			for _, contactGroup := range res.Data {
				if contactGroup.ID == rs.Primary.ID {
					return fmt.Errorf("contact group (%s) still exists", rs.Primary.ID)
				}
			}
		}
	}

	return nil
}

func testAccCheckContactGroupExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Missing Contact Group ID")
		}

		_, err := APIClient().GetContactGroup(context.TODO(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf(
				"cannot find contact group %s status %d errors %s %v: %w",
				rs.Primary.ID,
				err.(statuscake.APIError).Status,
				err.(statuscake.APIError).Message,
				err.(statuscake.APIError).Errors,
				err.(statuscake.APIError),
			)
		}

		return nil
	}
}

func TestAccContactGroup_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckContactGroupDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = "[TEST] Basic"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists("statuscake_contact_group.main"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "name", "[TEST] Basic"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "ping_url", ""),
				),
			},
		},
	})
}

func TestAccContactGroup_full(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckContactGroupDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name = "[TEST] Complete"
						ping_url = "https://ping.me/not"

						mobile_numbers = [
							"+49 176 66666666",
							"+49 176 55555555",
						]
						email_addresses = [
							"mb@mail.com",
							"another@email.com",
						]
						integration_ids = [
							"88888",
							"99999",
						]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists("statuscake_contact_group.main"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "name", "[TEST] Complete"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "ping_url", "https://ping.me/not"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "mobile_numbers.0", "+49 176 55555555"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "mobile_numbers.1", "+49 176 66666666"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "email_addresses.0", "another@email.com"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "email_addresses.1", "mb@mail.com"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "integration_ids.0", "88888"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "integration_ids.1", "99999"),
				),
			},
		},
	})
}

func TestAccContactGroup_validations(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckContactGroupDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = ""
					}
				`,
				ExpectError: regexp.MustCompile("expected \"name\" to not be an empty string"),
			},
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = "[TEST] Valid Name Invalid Url"
						ping_url = "this-is-not-valid"
					}
				`,
				ExpectError: regexp.MustCompile("expected \"ping_url\" to have a host"),
			},
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = "[TEST] Valid Name Invalid Url"
						ping_url = "http/this-is-also-not-valid"
					}
				`,
				ExpectError: regexp.MustCompile("expected \"ping_url\" to have a host"),
			},
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = "[TEST] Valid Name Invalid Protocol"
						ping_url = "ftp://ftp.google.com/invalid-protocol"
					}
				`,
				ExpectError: regexp.MustCompile("expected \"ping_url\" to have a url with schema of: \"http,https\""),
			},
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name     = "[TEST] All here is valid"
						ping_url = "https://www.example.com"

						mobile_numbers  = []
						email_addresses = []
						integration_ids = []
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists("statuscake_contact_group.main"),
				),
			},
		},
	})
}

func TestAccContactGroup_changing(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckContactGroupDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name = "[TEST] Complete"
						ping_url = "https://ping.me/not"

						mobile_numbers = [
							"+49 176 66666666",
							"+49 176 55555555",
						]
						email_addresses = [
							"mb@mail.com",
							"another@email.com",
						]
						integration_ids = [
							"88888",
							"99999",
						]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists("statuscake_contact_group.main"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "name", "[TEST] Complete"),
				),
			},
			{
				Config: `
					resource "statuscake_contact_group" "main" {
						name = "[TEST] Complete Updated V1"
						ping_url = "https://ping.me/not-v1"

						mobile_numbers = [
							"+49 176 99999999",
							"+49 176 88888888",
						]
						email_addresses = [
							"mb+v1@mail.com",
							"another+v1@email.com",
						]
						integration_ids = [
							"77777",
							"66666",
						]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists("statuscake_contact_group.main"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "name", "[TEST] Complete Updated V1"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "ping_url", "https://ping.me/not-v1"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "mobile_numbers.0", "+49 176 88888888"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "mobile_numbers.1", "+49 176 99999999"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "email_addresses.0", "another+v1@email.com"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "email_addresses.1", "mb+v1@mail.com"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "integration_ids.0", "66666"),
					resource.TestCheckResourceAttr("statuscake_contact_group.main", "integration_ids.1", "77777"),
				),
			},
		},
	})
}
