StatusCake Terraform Provider
==================
<p align="center" style="display: flex;justify-content: center; align-items: center; height: 200px;">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" height="100px">
</p>

Welcome to the Status Terraform Provider! This provider is heavily inspired by the official one but targeting the new functionalities of the [StatusCake V1.0-beta API](https://www.statuscake.com/api/v1/).

To view the full documentation of this provider, we recommend checking the [Terraform Registry](https://registry.terraform.io/providers/marceloboeira/statuscake/latest)!


Releases
---------

* [v1.0.0-rc2](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v1.0.0-rc2) - Bug fixes, stable release candidate.
* [v1.0.0-rc1](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v1.0.0-rc1) - First usable release, see feature matrix for more details.
* [v0.1.0](https://github.com/marceloboeira/terraform-provider-statuscake/releases/tag/v0.1.0) - Test release, making sure the whole process works - Non-Functional.

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

Contributing
------------

Read our [contributors](https://github.com/marceloboeira/terraform-provider-statuscake/docs/CONTRIBUTING.md) guide for more info on contributing.
