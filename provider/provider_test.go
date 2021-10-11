package provider

import (
	"os"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var providerFactory = map[string]func() (*schema.Provider, error){
	"statuscake": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func testAccPreCheck(t *testing.T) func() {
	return func() {
		t.Helper()

		if key := os.Getenv(ENV_APIKEY); key == "" {
			t.Fatal("Missing STATUSCAKE_APIKEY")
		}
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()

	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func APIClient() *statuscake.APIClient {
	return statuscake.NewAPIClient(os.Getenv(ENV_APIKEY))
}
