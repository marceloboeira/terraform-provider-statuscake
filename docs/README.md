StatusCake Terraform Provider
==================
<p align="center" style="display: flex;justify-content: center; align-items: center; height: 200px;">
    <img src="https://raw.githubusercontent.com/hashicorp/terraform-website/master/content/source/assets/images/logo-terraform-main.svg" height="100px">
</p>

Welcome to the Status Terraform Provider! This provider is heavily inspired by the official one but targeting the new functionalities of the [StatusCake V1.0-beta API](https://www.statuscake.com/api/v1/).

To view the full documentation of this provider, we recommend checking the [Terraform Registry](https://registry.terraform.io/providers/marceloboeira/statuscake/latest)!


Releases
---------

* [v1.0.0-rc3](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v1.0.0-rc4) - Bug fixes, tests, improved error messages
* [v1.0.0-rc2](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v1.0.0-rc2) - Bug fixes, stable release candidate.
* [v1.0.0-rc1](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v1.0.0-rc1) - First usable release, see feature matrix for more details.
* [v0.1.0](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v0.1.0) - Test release, making sure the whole process works - Non-Functional.

Usage
------

This is, for now, a complementary provider for [StatusCakeDev/statuscake](http://github.com/statusCakeDev/terraform-provider-statuscake). It supports mainly missing features, such as API V1 Contact Groups.

Therefore, usage is recommended side by side, like this:

```hcl
terraform {
  required_providers {
    statuscake = {
      source  = "StatusCakeDev/statuscake"
      version = "1.0.1"
    }
    # Important: Use an alternative name to avoid provider conflicts
    bolo = {
      source  = "marceloboeira/statuscake"
      version = "1.0.0-rc3"
    }
  }
}

provider "bolo" {
  # StatusCake API Key
  apikey = "...."
}

## Resource Names Continue with `statuscake_` namepsace
resource "statuscake_contact_group" "staff-engineering" {
  provider = bolo
  name     = "Staff Engineers"
  ping_url = "staf"

  mobile_numbers = [
    "+49 176 99999999",
    "+49 176 88888888",
  ]

  email_addresses = [
    "noreply@marceloboeira.com",
    "please.noreply@marceloboeira.com",
  ]

  integration_ids = [
    "111",
    "333",
    "555",
  ]
}
```

Feature Matrix
-------------

|        Entity | Feature | Status               |
|--------------:|---------|:--------------------:|
| Contact Group | Create  |  ✅ (>= v1.0.0-rc1)  |
|               | Update  |  ✅ (>= v1.0.0-rc1)  |
|               | Delete  |  ✅ (>= v1.0.0-rc1)  |
|               | Import  |  ✅ (>= v1.0.0-rc1)  |
| Test          | Create  |  ⛔️                  |
|               | Update  |  ⛔️                  |
|               | Delete  |  ⛔️                  |
|               | Import  |  ⛔️                  |
| Integrations  | Create  |  ⛔️                  |
|               | Update  |  ⛔️                  |
|               | Delete  |  ⛔️                  |
|               | Import  |  ⛔️                  |

* ✅ - Supported
* ⛔️ - Not supported

##### Details

* **Tests** - Not priority as the core provider supports this, PRs are welcome
* **Integrations** - Not possible since there is no API from StatusCake side to enable the terraform provider

### Troubleshooting

* StatusCake has a very low threshold on req/s, if you have many tests or contact groups you might want to limit the number of resources/s being manipulated by terraform. e.g.: `terraform plan -parallelism=1`.

Contributing
------------

Read our [contributors](https://github.com/marceloboeira/terraform-provider-statuscake/docs/CONTRIBUTING.md) guide for more info on contributing.
