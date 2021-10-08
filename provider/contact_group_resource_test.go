package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccContactGroup_create(t *testing.T) {
	tResource := "contact_group.example"
	gcName := "Core Infra"

	resource.Test(t, resource.TestCase{
		PreCheck:     testPreCheck,
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContactGroupBasic(gcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContactGroupExists(tResource),
					resource.TestCheckResourceAttr(tResource, "name", gcName),
					resource.TestCheckResourceAttr(tResource, "ping_url", ""),
					//resource.TestCheckResourceAttr(tResource, "mobile_numbers", []string{}),
				),
			},
		},
	})
}

func testAccCheckContactGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*statuscake.APIClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "contact_group" {
			continue
		}

		_, err := client.GetContactGroup(context.Background(), rs.Primary.ID).Execute()
		if err == nil {
			return fmt.Errorf("Alert still exists")
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccContactGroupBasic(n string) string {
	return fmt.Sprintf(`
resource "contact_group" "example" {
  name        = "%s"
}
`, n)
}

func testAccCheckContactGroupExists(r string) error {
	return func(state *terraform.State) error {
		rs := state.RootModule().Resources[r]
		if rs == nil {
			return fmt.Errorf("Not found: %s", r)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		name := rs.Primary.ID
		client := testAccProvider.Meta().(*statuscake.APIClient)
		_, err := client.GetContactGroup(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", r, err)
		}
		return nil
	}
}
