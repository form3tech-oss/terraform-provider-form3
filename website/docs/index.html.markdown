---
layout: "form3"
page_title: "Provider: Form3"
sidebar_current: "docs-form3-index"
description: |-
  The Form provider is used to interact with the resources supported by Form3. The provider needs to be configured with the proper client id and secret before it can be used.
---

# Form3 Provider

The Form3 provider is used to interact with the
resources supported by Form3. The provider needs to be configured
with a client id and secret before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Form3 provider
variable "api_host" {}
variable "client_id" {}
variable "client_secret" {}
variable "organisation_id" {}

provider "form3" {
  api_host      = "${var.api_host}"
  client_id     = "${var.client_id}"
  client_secret = "${var.client_secret}"
}


resource "form3_user" "admin_user" {
  organisation_id = "${var.organisation_id}"
  user_id = "44247ebb-fe01-44ab-887d-7f344481712f"
  user_name = "terraform-user"
  email = "terraform-user@form3.tech"
  roles = ["ad538853-4db0-44e3-9369-17eaae4aa3b7"]
}
```

## Argument Reference

The following provider arguments are supported:

* `api_host` - (Optional) The Form3 api host, this defaults to `api.form3.tech`.
This can also be specified with the `FORM3_HOST` shell
  environment variable.
* `client_id` - (Required) The Form3 client id used to access the api.
  This can also be specified with the `FORM3_CLIENT_ID` shell
  environment variable.
* `client_secret` - (Required) The Form3 client secret used to access the api.
  This can also be specified with the `FORM3_CLIENT_SECRET` shell
  environment variable.