---
layout: "form3"
page_title: "Form3: form3_user"
sidebar_current: "docs-form3-resource-user"
description: |-
  Provides a Form3 user resource.
---

# form3\_user

A [user](http://api-docs.form3.tech/#user-resource) resource.

## Example Usage

```hcl
# Add a user to your form3 account
resource "form3_user" "admin_user" {
  organisation_id = "${var.organisation_id}"
  user_id = "44247ebb-fe01-44ab-887d-7f344481712f"
  user_name = "terraform-user"
  email = "terraform-user@form3.tech"
  roles = ["ad538853-4db0-44e3-9369-17eaae4aa3b7"]
}
```

## Argument Reference

The following arguments are supported:

* `organisation_id` - (String, Required) The origanisation id to create this user in.
* `user_id` - (String, Required) Unique identifier for this user.
* `user_name` - (String, Required) User name for this user.
* `email` - (String, Required) User's email address.
* `roles` - (List) A list of roles id's that should be assigned to this user.


## Attributes Reference
The following attributes are exported:

* `id` - The ID of the user.