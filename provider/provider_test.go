package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProviders = map[string]*schema.Provider{
		"statuscake": testAccProvider,
	}
	testAccProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testPreCheck() {
	return func(t *testing.T) {
		if k := os.Getenv("STATUSCAKE_APIKEY"); k == "" {
			t.Fatalf("Missing required environment variable: %s", "STATUSCAKE_APIKEY")
		}
	}
}
