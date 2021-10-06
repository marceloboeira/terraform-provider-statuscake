terraform {
  required_version = ">= 1.0.0"
  required_providers {
    mb = {
      source  = "marceloboeira/statuscake"
      version = "0.1.0"
    }
  }
}

variable "statuscake_api_key" {
  description = "The StatusCake API key"
}

provider "mb" {
  apikey = var.statuscake_api_key
}

resource "statuscake_contact_group" "my_fancy_group" {
  provider = mb
  name     = "Call Me Maybe"
  ping_url = "https://marceloboeira.com/dont-touch-this"

  mobile_numbers = [
    "+49 176 99999999",
    "+49 176 88888888",
    "+49 176 77777777",
  ]

  email_addresses = [
    "noreply@marceloboeira.com",
    "please.noreply@marceloboeira.com",
  ]

  integration_ids = [
    "99999",
    "88888",
  ]
}
