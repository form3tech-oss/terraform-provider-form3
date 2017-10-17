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
provider "runscope" {
  access_token = "${var.access_token}"
}

# Create a bucket
resource "runscope_bucket" "main" {
  name         = "terraform-ftw"
  team_uuid    = "870ed937-bc6e-4d8b-a9a5-d7f9f2412fa3"
}

# Create a test in the bucket
resource "runscope_test" "api" {
  name         = "api-test"
  description  = "checks the api is up and running"
  bucket_id    = "${runscope_bucket.main}"
}
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Required) The Form3 access token.
  This can also be specified with the `RUNSCOPE_ACCESS_TOKEN` shell
  environment variable.
* `api_url` - (Optional) If set, specifies the Form3 api url, this
   defaults to `"https://api.runscope.com`. This can also be specified
   with the `RUNSCOPE_API_URL` shell environment variable.
